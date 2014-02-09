goog.provide('glosure');

goog.require('pkg');

glosure.someFunc = function() {
  pkg.someFunc();
}

glosure.someFunc();

