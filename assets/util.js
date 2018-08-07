'use strict'

const axios = require('axios')

const myAxios = axios.create({
    baseURL: location.href.replace(/:[0-9].*/, ':1312'),
    withCredentials: true
})

const getSAMLAPI = async () => {
    const res = await myAxios.get('/util/get-config/SAMLAPI.url')

    return res.data
}

(async () => {
    exports.axios = myAxios

    exports.axiosSAML = axios.create({
        baseURL: await getSAMLAPI()
    })
})()