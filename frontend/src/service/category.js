'use strict';

const request = require('request');

function CategoryService(url){
  this.url = url;
}

CategoryService.prototype.getById = function (id, cb) {
  request({
    url: this.url+'/categories/'+id,
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  }, (err, response, body) => {
    if(err || response.statusCode != '200'){
      cb("Cannot Get Me", null);
    }else{
      cb(null, body)
    }
  });
};

CategoryService.prototype.getAll = function (cb) {
  request({
    url: this.url+'/categories',
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  }, (err, response, body) => {
    if(err || response.statusCode != '200'){
      cb("Cannot Get Me", null);
    }else{
      cb(null, body)
    }
  });
};

module.exports = CategoryService;
