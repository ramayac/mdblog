// Apply saved theme immediately (loaded in <head>) to prevent flash of wrong theme.
(function () {
    function readTheme() {
        try { return localStorage.getItem('theme-mode'); } catch (_) { return null; }
    }
    function saveTheme(mode) {
        try { localStorage.setItem('theme-mode', mode); } catch (_) {}
    }

    var saved = readTheme();
    if (saved) {
        document.documentElement.setAttribute('data-theme', saved);
    }

    function toggleTheme() {
        var current = document.documentElement.getAttribute('data-theme');
        if (!current) {
            current = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
        }
        var next = current === 'dark' ? 'light' : 'dark';
        document.documentElement.setAttribute('data-theme', next);
        saveTheme(next);
    }

    function closeDropdowns(except) {
        document.querySelectorAll('details.nav-dropdown[open]').forEach(function (d) {
            if (d !== except) d.removeAttribute('open');
        });
    }

    document.addEventListener('DOMContentLoaded', function () {
        var btn = document.querySelector('.theme-toggle');
        if (btn) btn.addEventListener('click', toggleTheme);
    });

    document.addEventListener('click', function (e) {
        closeDropdowns(e.target.closest('details.nav-dropdown'));
    });

    document.addEventListener('keydown', function (e) {
        if (e.key === 'Escape') closeDropdowns(null);
    });
})();
