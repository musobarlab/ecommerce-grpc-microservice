'use strict';

const request = require('request');

function MeService(url){
  this.url = url;
}

MeService.prototype.me = function (token, cb) {
  request({
    url: this.url+'/me',
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Authorization': token,
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

module.exports = MeService;
