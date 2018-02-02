'use strict';

const LocalStorage = require('node-localstorage').LocalStorage;

function TokenStorage(){
  this.localStorage;
  if (typeof localStorage === "undefined" || localStorage === null) {
    this.localStorage = new LocalStorage('./scratch');
  }
}

TokenStorage.prototype.store = function(token){
  this.localStorage.setItem('token', token);
}

TokenStorage.prototype.get = function(){
  return this.localStorage.getItem('token');
}

TokenStorage.prototype.remove = function(){
  return this.localStorage.removeItem('token');
}

module.exports = TokenStorage;
