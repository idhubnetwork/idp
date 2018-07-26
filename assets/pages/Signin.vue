<template>
    <div>
        <h1 class="title">{{ $t('signin') }}</h1>
        <h1 class="subtitle">
            <textarea cols="20" rows="5" class="textarea" type="text" v-model="message"></textarea>
            <button class="button is-link is-rounded" @click="sign" :disabled="validate">{{ $t('signin') }}</button>
        </h1>
        <h1 class="subtitle">
            <textarea class="textarea" type="text" v-model="sigStringify" readonly></textarea>
        </h1>
    </div>
</template>

<script>
import { axios } from '@/util'

export default {
    props: [
        'coinbase'
    ], data() {
        return {
            message: '',
            signature: {}
        }
    }, mounted() {
        if (this.coinbase !== '' && typeof this.coinbase != 'undefined') {
            this.getBookingMsg()
        }
    }, computed: {
        validate() {
            if (this.message === '') return true
            return this.coinbase === null || this.coinbase === ''
        }, sigStringify() {
            return JSON.stringify(this.signature, '', 4)
        }
    }, methods: {
        sign() {
            web3.personal.sign(web3.fromUtf8(this.message), this.coinbase, (err, res) => {
                if (err) this.signature = err
                if (res) {
                    this.signature = {
                        address: this.coinbase,
                        msg: this.message,
                        sig: res,
                        version: "3",
                        signer: "web3"
                    }

                    axios.post('/auth/verify', {
                        addr: this.coinbase,
                        sig: this.signature.sig
                    }).then(res => {
                        this.signature = Object.assign({}, this.signature, { verified: 'OK' })
                        location.href = '/'
                    }).catch(err => {
                        this.signature = err.response
                        location.href = '/'
                    })
                }
            })
        },
        getBookingMsg() {
            axios.post('/auth/booking', {
                addr: this.coinbase
            }).then(res =>
                this.message = res.data
            ).catch(
                err => console.log(err)
            )
        }
    }, watch: {
        coinbase() {
            this.getBookingMsg()
        }
    }
}
</script>
