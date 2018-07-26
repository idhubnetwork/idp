'use strict'

const axios = require('axios')

exports.axios = axios.create({
    baseURL: location.href.replace(/:[0-9].*/,':1312'),
    withCredentials: true
})

exports.axiosSAML = axios.create({
    baseURL: 'http://58.83.219.136:33333'
})