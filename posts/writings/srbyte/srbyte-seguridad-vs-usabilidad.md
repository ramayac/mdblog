---
title: Seguridad vs. Usabilidad
date: 2010-08-05
author: Robertux
tags: programacion, usabilidad, web, seguridad, diseño
draft: false
post_id: blog-3515952828243908885.post-3380693066006441941
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiDhVaQFgONzIRKH0_OYUjc1oxvrKf0zBGzKWBJ3ZjWfzp9Bm91rUWLdgqr74xyFhffHAS8iAaw-KZSAjeuCwSwdj6VAv0-qb-dmGZsitLSkvchsXpK8VV2f6hQes_KJJuHa9ajLkNXjOQ/s400/BancoAgricolaFAIL.png)    
"La imagen de la discordia, origen y razón de ser de este post"

Recientemente por casualidades del destino me ví en la necesidad de ingresar al sitio en línea del banco agrícola. Tenía meses de no entrar a dicho sitio y al ingresar en esta ocasión me doy cuenta de un cambio significativo que han aplicado en la forma como escribes tu usuario, clave y token para ingresar al sistema. Resulta que ahora ya no puedes digitar de forma directa estos datos sino que debes hacer uso de un teclado virtual que te aparece en pantalla. Tal como se aprecia en la imagen inicial de este post.

Quizá la característica más particular que he encontrado junto con otros usuarios que pasaron por la misma experiencia de tratar de hacer login en dicha página, ha sido la que las teclas no se encuentran en un orden aparente sino mas bien ordenadas aleatoriamente. De hecho, por cada vez que refrescas la página (ej. en un intento fallido), dichas teclas vuelven a cambiar de posición y nuevamente de manera aleatoria.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhViu6z5nJNsR_osQHJoC8TXPNNUbC7uiV8wlMJKzvvN_rX0ZP70tvzuJP0FKx-RggQdHv17oIeYcPk-WDGx52uiMqXdMQ3IcORlqWisdCgE9LTW_f5rjQayw-CYJLBrdSv117xFdJgpjg/s400/BancoAgricolaFAILTeclado.png)    
"A ver... hoy donde quedó la A? y la L? la Ñ? OMG!"

Es cierto que hoy en día las instituciones bancarias sufren ataques de todo tipo en sus sitios web, el phishing y el cross site scripting por ejemplo, con el cual cualquier otro intenta vulnerar los sistemas web para tener acceso a las cuentas de sus clientes y sus datos privados están en juego. Es cierto que es deber de cada empresa en general proteger los datos de sus usuarios especialmente en el mayor punto débil de todos los sistemas seguros: sus pantallas de ingreso.

Todo esto es necesario para asegurarnos que los sitios a los que ingresamos se encuentren seguros de cualquier ataque pero no por ello se debe dejar de lado la usabilidad del sitio y la experiencia del usuario. Es posible mantener un balance entre seguridad y usabilidad?

La respuesta es SI.

Asumiendo que la inclusión de este terrible teclado virtual con teclas aleatorias que les mencionaba al principio se incluyó para prevenir que se intente ingresar por medio de un programa y no mediante la interacción con un usuario, existen alternativas mucho más usables para lograr este fin. Menciono las dos más comunes y útiles:

CAPTCHA

Captcha es un sistema ideado con la intención de diferenciar si la interacción con un sitio web la está realizando un humano o un programa de computadora. Para ello, muestran una imagen la cual contiene una palabra con el texto un tanto distorsionado o difuso, de manera que solamente un humano pueda leer e interpretar el texto, para luego escribirlo en un campo al final del formulario. No es 100% infalible ya que hoy en día existen programas para la interpretación de texto difuso pero sí es bastante confiable, ya que tanto Google, Facebook y otros sitios importantes hacen uso del mismo. Uno de los CAPTCHAs más populares es [reCAPTCHA](https://www.google.com/recaptcha) de Google, el cual es usado en la mayoría de formularios de login de Google cuando deseas crear una cuenta o cuando te has equivocado más de tres veces en tu login:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhgcc6oiRWfOmFeH-ivTorzAaile1SvWDyQNjDmkxh5FHUuDHq332buXzYA4DjDMAvdyV_Q7cqB7XoIU-XcGUXVHrPSfuKWe9pvE1K-TEr46-WQzjTlcCP42EODpP0QpKieBclwCRLdSRo/s400/ReCAPTCHA.png)    
"Ejemplo de un cuadro de reCAPTCHA. Consiste en escribir el texto difuso en la caja de texto. En caso de no distinguirlo, click en el botón con forma de bocina para escuchar su pronunciación"

En [este link](https://www.google.com/recaptcha/whyrecaptcha) puedes obtener plugins para blogs y el código fuente para hacer uso de reCAPTCHA en tu sitio.

Slide to Unlock

Este es un muy ingenioso sistema de ingreso que originalmente viene en los iPhone/iPod Touch, el cual se puede usar también en los sitios web para asegurarse que es un humano quien hace clic sobre el botón y lo desliza de izquierda a derecha. Esta interacción sería muy difícil (mas no imposible) de llevarse a cabo mediante un programa que intentara llenar automáticamente un formulario de login por lo que se podría considerar efectivo y fácil de usar.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiPUDFTaHLrEFVBQpb9lNjSxiTnSDcRVPHQs3Dpq7zCaZuI8HE0fYvBOrsG-ZG_vZ2K83KBRq1yoSsdkSBX_cu2ToZEbsNaNw0gQsHbm74ueplaD9ZPzsWDIZZPdNzIKOe8VY1eN9tBytU/s400/slidecaptcha-412x300.jpg)    
"Ejemplo del mecanismo de submit 'Slide to Submit' similar al 'Slide to Unlock' del iPhone/iPod Touch"

En [este link](https://www.aboone.com/javascript-iphone-lock-slider-with-jquery) puedes encontrar el código fuente para incluir un botón de tipo "Slide to submit" mediante código Javascript.

Para mi punto de vista, usando alguna de las técnicas aquí presentadas, sumándole el uso de conexiones seguras por HTTPS, tokens, certificados digitales y mecanismos de prevención del XSS, un sitio podría considerarse lo suficientemente seguro como para no tener la necesidad de molestar más al usuario con pasos extra o peticiones para llevar a cabo su ingreso al sitio en cuestión.

Alguna otra técnica o sugerencia para crear sitios web tanto seguros como usables?

Bonus: Talvez estos [91 ejemplos de cajas de login](https://www.smileycat.com/design_elements/login_forms/) te puedan inspirar para crear una página completamente usable y agradable a la vista de tus usuarios.

Bonus No. 2: Me entero en twitter que [@hkadejo](https://twitter.com/hkadejo) desarrolló un muy útil bookmarklet para desactivar las cajas de texto y poder escribir de forma manual el usuario y clave en la página del Banco Agrícola. Pueden encontrarlo en [este link](https://www.hkadejo.com/?p=48).