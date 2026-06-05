---
title: Herramientas Online de Formato e Indentación de Código
date: 2010-04-21
author: Robertux
tags: codigo, css, herramienta, java, xml, json, online, html, javascript, sql
draft: false
post_id: blog-3515952828243908885.post-4503273520819917409
---

Mas de alguna vez nos ha tocado revisar algún mensaje XML o JSON fuera de la oficina, donde no tenemos nuestras herramientas de desarrollo que tan agradablemente se encargan de formatear e indentar nuestros archivos para que podamos leerlos, analizarlos y encontrar posibles fallas en ellos. En estas ocasiones nos ha tocado manualmente estar partiendo el archivo en diferentes líneas y agregando sangrías para darle el formato adecuado y se nos vuelva más fácil su análisis

Si de casualidad nos encontramos en estas circunstancias pero tienen la bondad de contar con una conexión a Internet, pueden aprovechar algunas herramientas online que se encargan de formatear, colorear e indentar nuestros archivos de código fuente.

Hay muchos otros que no he listado acá por el hecho de que solamente colorean el código y considero mas útiles aquellos que además del syntax highlight también te agregan los respectivos espacios de tabulación.

XML:

- [https://xmlindent.com/](https://xmlindent.com/)
- [https://x01.co.uk/tools/online-xml-formatter](https://x01.co.uk/tools/online-xml-formatter)

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhlrliqe46s7W3YHnlzL-BHfnhifbn4Ne19YLU2kMNQrCtS_pF6M1PBTxvLSXzOVWaEjr0JPmIt3NB4-UPQJtPh7hRyqD-eloxvucbG8VH4Rs4fmGyrywm21ohnAxbNoEdCypuptQGjYNQ/s400/XMLIndent.png)    
"Captura de pantalla de xmlindent.com, formateando un xml de ejemplo (click para agrandar)"

JSON:

- [https://jsoneditor.appspot.com/](https://jsoneditor.appspot.com/)
- [https://jsoneditor.net/](https://jsoneditor.net/)

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj1LAxvPxeT29INjShraJL1S2VPh0W8TJy0-x3xpP0TVRvzeAOJqSRvjm2VbHFcKc5bVSgSA_KXpxmyCmW_r0mSlwzUVS7PgE0f9dn-nFk45DW_hiqjy2px9Nh8i7tnT9ZlGOvjzrKr62Q/s400/JSONEditor.png)    
"Captura de pantalla de jsoneditor.net, quien te forma la estructura jerárquica de tu mensaje JSON"

Múltiples lenguajes de programación:

- [https://infohound.net/tidy/](https://infohound.net/tidy/) (HTML)
- [https://www.prettyprinter.de/](https://www.prettyprinter.de/) (PHP, Java, C++, C, Perl, Javascript, CSS)
- [https://www.dpriver.com/pp/sqlformat.htm](https://www.dpriver.com/pp/sqlformat.htm) (SQL, PL/SQL)
- [https://jsbeautifier.org/](https://jsbeautifier.org/) (Javascript)
- [https://tools.arantius.com/tabifier](https://tools.arantius.com/tabifier) (html/css)

Cabe resaltar algunas características especiales que poseen algunas de estas herramientas como por ejemplo jsbeautifier además de formatear el código, también te lo desempaqueta cuando este ha sido generado mediante la herramienta [packer](https://www.blogger.com/dean.edwards.name/packer/), o la herramienta sqlformat que te genera un output de SQL no solamente formateado y con sintaxis coloreada sino que también te puede generar un output de SQL que generalmente colocamos en una variable String dentro de nuestro lenguaje de programación, así como lo ilustra la siguiente imagen:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjtUGcpHfq903t3q_k1eStBLb072UkMWmiOL1DX_F_HYztZchricXgUHlJvz2bw2MdIFtgrBEe67EAAiL3ut0BrvI_upy59PJS5ZHsUmRVk5J6bZoksCr2jOsVk9tELDcJdfFFaWIT1Z3c/s400/SqlFormat.png)    
"Sqlformat, indentando correctamente un procedimiento almacenado generando un output a manera de StringBuffer para ser utilizado en una clase Java (Click para agrandar)"

Algún otro que ustedes utilicen para formatear su código fuente en línea?