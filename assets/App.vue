<template>
    <div>
        <NavBar :coinbase="coinbase" :identity="identity"></NavBar>

        <section class="section">
            <div class="container is-fluid">
                <router-view :coinbase="coinbase" :identity="identity" :date="date" ></router-view>
            </div>
        </section>

        <FooterBar></FooterBar>
    </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import FooterBar from '@/components/FooterBar.vue'

export default {
    components: {
        NavBar, FooterBar
    }, data() {
        return {
            date: new Date().toLocaleString(),
            identity: document.cookie.split(';').filter(c => c.startsWith('IDHUB_IDENTITY')).join('').split('=')[1],
            coinbase: ''
        }
    }, mounted() {
        setInterval(() => {
            this.coinbase = typeof web3 === 'undefined' || web3.eth.coinbase === null ? '' : web3.eth.coinbase
        }, 1000)
    }
}
</script>

<style lang="scss" scoped>
div {
  outline: none;
}
</style>

<style lang="scss">
h1 {
  &.title {
    color: #333f55;
    font-size: 36px;
  }
}

.control {
  input {
    background-color: #f6f6f6;
  }
}

.section {
  background-color: #f2f4f6;
}

.columns {
  background: #ffffff;
  padding: 30px 0;
  margin-top: 10px;
}
</style>
