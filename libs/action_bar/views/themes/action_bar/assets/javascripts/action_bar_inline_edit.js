"use strict";var _typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t};!function(t){"function"==typeof define&&define.amd?define(["jquery"],t):t("object"===("undefined"==typeof exports?"undefined":_typeof(exports))?require("jquery"):jQuery)}(function(t){function n(o,e){this.$element=t(o),this.options=t.extend({},n.DEFAULTS,t.isPlainObject(e)&&e),this.init()}var o,e=t("body"),i="qor.actionbar.inlineEdit",r="enable."+i,a="disable."+i,c="click."+i,u=".qor-actionbar-button[data-url]",d="qor-actionbar-iframe";return n.prototype={constructor:n,init:function(){this.bind(),this.initStatus()},bind:function(){this.$element.on(c,u,this.click),t(document).on("keyup",this.keyup)},initStatus:function(){var n=document.createElement("iframe");n.src=o,n.id=d,n.attachEvent?n.attachEvent("onload",function(){t(u).show()}):n.onload=function(){t(u).show()},document.body.appendChild(n)},keyup:function(t){var n=document.getElementById("qor-actionbar-iframe");27===t.keyCode&&n&&n.contentDocument.querySelector(".qor-slideout__close").click()},click:function(){var n=t(this),o=n.data(),i=document.getElementById("qor-actionbar-iframe"),r=i.contentWindow.$,a=r("body").data("qor.slideout");if(a)return i.classList.add("show"),a.open(o),e.addClass("open-actionbar-editor"),!1}},n.plugin=function(o){return this.each(function(){var e,r=t(this),a=r.data(i);if(!a){if(/destroy/.test(o))return;r.data(i,a=new n(this,o))}"string"==typeof o&&t.isFunction(e=a[o])&&e.apply(a)})},t(function(){e.attr("data-toggle","qor.actionbars");var i='[data-toggle="qor.actionbars"]';o=t(u).data("iframe-url"),t(document).on(a,function(o){n.plugin.call(t(i,o.target),"destroy")}).on(r,function(o){n.plugin.call(t(i,o.target))}).triggerHandler(r)}),n});