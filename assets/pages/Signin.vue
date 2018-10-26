<template>
    <div>
        <h1 class="title">{{ $t('signin') }}</h1>
        <div>
            <h2> 登录流程提示 {{ date }} </h2>
            <li> 输入 DID </li>
            <li> 启动 Metamask </li>
            <li> 获取安全码 </li>
            <li> 签名登录 </li>
        </div>
        <h1 class="subtitle">
            <textarea cols="20" rows="5" class="textarea" type="text" v-model="identity"></textarea>
            <button class="button is-link is-rounded" @click="getBookingMsg" :disabled="identityValidate">{{ $t('get booking massage') }}</button>
        </h1>
        <h1 class="subtitle">
            <textarea cols="20" rows="5" class="textarea" type="text" v-model="message"></textarea>
            <button class="button is-link is-rounded" @click="sign" :disabled="messageValidate">{{ $t('signin') }}</button>
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
        'coinbase',
        'identity',
        'date'
    ], data() {
        return {
            message: '',
            signature: {}
        }
    }, mounted() {
        // if (this.coinbase !== '' && typeof this.coinbase != 'undefined') {
        //     this.getBookingMsg()
        // }
    }, computed: {
        identityValidate() {
            if (this.identity === '' || this.identity === 'undefined') {
                console.log('here' + this.identity)
                return true
            }
            return this.coinbase === null || this.coinbase === ''
        }, messageValidate() {
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
                        // address: this.coinbase,
                        address: this.identity,
                        msg: this.message,
                        sig: res,
                        version: "3",
                        signer: "web3"
                    }

                    axios.post('/auth/verify', {
                        // addr: this.coinbase,
                        addr: this.identity,
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
            console.log('here2, send booking massage\n' + 'identity is ' + this.identity + '\nCoinbase is ' + this.coinbase)
            axios.post('/auth/booking', {
                // addr: this.coinbase
                addr: this.identity
            }).then(res => {
                this.message = res.data
                console.log('here is the message' + this.message)
            }).catch(
                err => console.log(err)
            )
        }
    }, watch: {
        // coinbase() {
        //     this.getBookingMsg()
        // }
    }
}
</script>
