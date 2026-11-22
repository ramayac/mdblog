#!/usr/bin/env node
// Deterministic computed-geometry & runtime layout checker for MDBlog.
// Usage: node scripts/lint-layout.mjs [baseURL]
// Requires Playwright (e.g. `npx --yes playwright@latest` or CI setup).

const BASE_URL = process.argv[2] || 'http://localhost:8091';

let chromium;
try {
  ({ chromium } = await import('playwright'));
} catch {
  console.error(
    'Error: playwright package not found.\n' +
    'Install with: npm install -g playwright && npx playwright install chromium\n' +
    'Or run with NODE_PATH pointing to a Playwright installation.'
  );
  process.exit(1);
}

const ROUTES = [
  { path: '/', expectStatus: 200, kind: 'home' },
  { path: '/pages/about', expectStatus: 200, kind: 'page' },
  { path: '/content/writings/personal/', expectStatus: 200, kind: 'category' },
  { path: '/content/writings/substack/well-well-well', expectStatus: 200, kind: 'post' },
  { path: '/content/writings/srbyte/srbyte-cambios-en-los-terminos-de-uso-de-google-shared-en', expectStatus: 200, kind: 'post' },
  { path: '/content/writings/unalistaaparte/post-semana-3-ajedrez-blake-y-briboner-a', expectStatus: 200, kind: 'post' },
  { path: '/?q=linux', expectStatus: 200, kind: 'search' },
  { path: '/feed', expectStatus: 200, kind: 'feed' },
  { path: '/content/does-not-exist', expectStatus: 404, kind: '404' },
];

const VIEWPORTS = [
  { name: 'desktop', width: 1280, height: 800 },
  { name: 'mobile', width: 390, height: 844 },
];

const THEMES = ['light', 'dark'];

const failures = [];
let checksRun = 0;

const browser = await chromium.launch({ headless: true });

try {
  for (const vp of VIEWPORTS) {
    const context = await browser.newContext({
      viewport: { width: vp.width, height: vp.height },
    });

    // Track light vs dark background on / to confirm theme toggle works
    const rootBgByTheme = {};

    for (const theme of THEMES) {
      await context.addInitScript((mode) => {
        try { localStorage.setItem('theme-mode', mode); } catch (_) {}
      }, theme);

      for (const route of ROUTES) {
        const page = await context.newPage();
        const label = `[${vp.name}/${theme}] ${route.path}`;
        const consoleErrors = [];
        const failedAssets = [];
        const cspViolations = [];

        page.on('console', (msg) => {
          if (msg.type() === 'error') consoleErrors.push(msg.text());
        });
        page.on('pageerror', (err) => consoleErrors.push(err.message));
        page.on('response', (resp) => {
          const url = resp.url();
          if (url.startsWith(BASE_URL + '/assets/') && resp.status() >= 400) {
            failedAssets.push(`${resp.status()} ${url}`);
          }
        });

        await page.addInitScript(() => {
          window.__cspViolations = [];
          document.addEventListener('securitypolicyviolation', (e) => {
            window.__cspViolations.push(`${e.violatedDirective}: ${e.blockedURI}`);
          });
        });

        const resp = await page.goto(BASE_URL + route.path, { waitUntil: 'domcontentloaded' });
        checksRun++;

        if (!resp || resp.status() !== route.expectStatus) {
          failures.push(`${label}: HTTP status = ${resp?.status()}, want ${route.expectStatus}`);
          await page.close();
          continue;
        }

        const metrics = await page.evaluate(() => {
          const rect = (sel) => {
            const el = document.querySelector(sel);
            if (!el) return null;
            const r = el.getBoundingClientRect();
            return { left: r.left, right: r.right, top: r.top, bottom: r.bottom, width: r.width };
          };

          const bodyStyle = getComputedStyle(document.body);
          const contentEl = document.querySelector('.post-content');
          const contentStyle = contentEl ? getComputedStyle(contentEl) : null;
          const firstChildEl = contentEl ? contentEl.firstElementChild : null;
          const firstChildMarginTop = firstChildEl ? parseFloat(getComputedStyle(firstChildEl).marginTop) : null;
          const h1El = document.querySelector('h1');
          const h1Size = h1El ? parseFloat(getComputedStyle(h1El).fontSize) : null;

          return {
            scrollWidth: document.documentElement.scrollWidth,
            innerWidth: window.innerWidth,
            bodyBg: bodyStyle.backgroundColor,
            cspViolations: window.__cspViolations || [],
            menu: rect('.site-menu'),
            footer: rect('.site-footer'),
            header: rect('.post-header') || rect('.blog-title') || rect('.category-title'),
            postNav: rect('.post-navigation'),
            contentRect: rect('.post-content'),
            contentFontSize: contentStyle ? parseFloat(contentStyle.fontSize) : null,
            contentLineHeight: contentStyle ? parseFloat(contentStyle.lineHeight) : null,
            firstChildMarginTop,
            h1Size,
          };
        });

        if (route.path === '/') {
          rootBgByTheme[theme] = metrics.bodyBg;
        }

        // 1. Runtime cleanliness
        if (consoleErrors.length > 0) {
          failures.push(`${label}: console errors: ${consoleErrors.join('; ')}`);
        }
        if (failedAssets.length > 0) {
          failures.push(`${label}: failed assets: ${failedAssets.join('; ')}`);
        }
        if (metrics.cspViolations.length > 0) {
          failures.push(`${label}: CSP violations: ${metrics.cspViolations.join('; ')}`);
        }

        // 2. Zero horizontal overflow
        if (metrics.scrollWidth > metrics.innerWidth + 1) {
          failures.push(`${label}: horizontal overflow (scrollWidth=${metrics.scrollWidth} > innerWidth=${metrics.innerWidth})`);
        }

        // 3. Column alignment (left & right flush with .site-menu within 2px)
        if (metrics.menu) {
          for (const [name, box] of [
            ['header', metrics.header],
            ['post-content', metrics.contentRect],
            ['post-navigation', metrics.postNav],
            ['footer', metrics.footer],
          ]) {
            if (!box) continue;
            if (Math.abs(box.left - metrics.menu.left) > 2) {
              failures.push(`${label}: ${name} left (${box.left.toFixed(1)}) misaligned with menu left (${metrics.menu.left.toFixed(1)})`);
            }
            if (Math.abs(box.right - metrics.menu.right) > 2) {
              failures.push(`${label}: ${name} right (${box.right.toFixed(1)}) misaligned with menu right (${metrics.menu.right.toFixed(1)})`);
            }
          }

          // 4. Vertical rhythm from nav bottom to primary page header
          if (metrics.header) {
            const gap = metrics.header.top - metrics.menu.bottom;
            if (gap < 12 || gap > 56) {
              failures.push(`${label}: nav-to-header vertical gap = ${gap.toFixed(1)}px (expected 12..56px)`);
            }
          }
        }

        // 5. Post/page typography & first-child margin bounds
        if (metrics.contentFontSize !== null) {
          if (metrics.contentFontSize < 15 || metrics.contentFontSize > 19) {
            failures.push(`${label}: .post-content font-size = ${metrics.contentFontSize}px (expected 15..19px)`);
          }
          if (metrics.contentLineHeight !== null) {
            const ratio = metrics.contentLineHeight / metrics.contentFontSize;
            if (ratio < 1.5 || ratio > 1.9) {
              failures.push(`${label}: .post-content line-height ratio = ${ratio.toFixed(2)} (expected 1.5..1.9)`);
            }
          }
          if (metrics.firstChildMarginTop !== null && metrics.firstChildMarginTop > 2) {
            failures.push(`${label}: .post-content > :first-child margin-top = ${metrics.firstChildMarginTop}px (expected <= 2px)`);
          }
          if (metrics.h1Size !== null && metrics.h1Size <= metrics.contentFontSize) {
            failures.push(`${label}: h1 (${metrics.h1Size}px) should be larger than .post-content (${metrics.contentFontSize}px)`);
          }
        }

        await page.close();
      }
    }

    if (rootBgByTheme.light && rootBgByTheme.dark && rootBgByTheme.light === rootBgByTheme.dark) {
      failures.push(`[${vp.name}]: light and dark themes rendered identical body background (${rootBgByTheme.light})`);
    }

    await context.close();
  }
} finally {
  await browser.close();
}

if (failures.length > 0) {
  console.error(`Layout lint FAILED (${failures.length} issues across ${checksRun} page checks):\n`);
  for (const f of failures) console.error('  - ' + f);
  process.exit(1);
}

console.log(`Layout lint PASSED (${checksRun} page/viewport/theme checks, 0 issues).`);
