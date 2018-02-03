'use strict';

const request = require('request');

function ProductService(url){
  this.url = url;
}

ProductService.prototype.getById = function (id, cb) {
  request({
    url: this.url+'/products/'+id,
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

ProductService.prototype.getAll = function (cb) {
  request({
    url: this.url+'/products',
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

module.exports = ProductService;
