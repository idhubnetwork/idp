<template>
    <div>
        <h1 class="title">{{ $t('metadata-manager') }}</h1>
        <h1 class="subtitle">
            <button class="button is-link is-rounded" @click="query">{{ $t('query') }}</button>
            <a class="button is-link is-rounded" @click="download" ref="download" :disabled="metadata === ''">{{ $t('download-metadata') }}</a>
            <button class="button is-danger is-rounded" @click="register">{{ $t('register') }}</button>
        </h1>
        <h1 class="subtitle">
            <div class="control">
                <textarea class="textarea" type="text" v-model="metadata" rows="10" readonly></textarea>
            </div>
        </h1>
    </div>
</template>

<script>
import { axiosSAML as axios } from '@/util'

export default {
    props: [
        'identity'
    ], data() {
        return {
            metadata: ''
        }
    }, methods: {
        query() {
            axios.get(`/getMetadata/${this.identity}`).then(res =>
                this.metadata = res.data
            ).catch(err =>
                this.metadata = JSON.stringify(err)
            )
        }, register() {
            axios.get(`/setOrganizations/${this.identity}`).then(res =>
                this.metadata = JSON.stringify(res.data)
            ).catch(err =>
                this.metadata = JSON.stringify(err)
            )
        }, download(e) {
            if (this.metadata === '') {
                e.preventDefault()
            }

            const xml = new Blob([this.metadata], { type: 'text/xml' })

            this.$refs.download.href = URL.createObjectURL(xml)
            this.$refs.download.download = `${this.identity}.xml`
        }
    }
}
</script>
