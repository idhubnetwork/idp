<template>
    <div>
        <p class="title">{{ $t('auth-manager-list') }}</p>

        <table class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth">
            <thead>
                <tr>
                    <th>#</th>
                    <th>org. user</th>
                    <th>roles</th>
                    <th>created at</th>
                    <th>updated at</th>
                    <th>action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(user, index) in users" :key="user.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ user.orgUser }}</td>
                    <td>{{ user.roles }}</td>
                    <td>{{ user.createdAt }}</td>
                    <td>{{ user.updatedAt }}</td>
                    <td>
                        <button class="button is-link is-rounded is-small" @click="getSaml(user.id)">Get SAML</button>
                    </td>
                </tr>
            </tbody>
        </table>

        <form action="https://signin.aws.amazon.com/saml" method="post" target="_blank">
            <button class="button is-warning is-rounded" type="submit" :disabled="saml === ''">AWS I'm comming...</button>
            <button class="button is-link is-rounded" type="button" @click="saml = ''">Clear</button>
            <br>
            <br>
            <div class="control">
                <textarea class="textarea" type="text" v-model="saml" rows="10" name="SAMLResponse" readonly></textarea>
            </div>
        </form>
    </div>
</template>

<script>
import { axios } from '@/util'

export default {
    data() {
        return {
            users: [],
            saml: ''
        }
    }, mounted() {
        axios.get(`/auth-manager/authorized-list`).then(res => {
            this.users = res.data
        }).catch(err =>
            console.log(err.response)
        )
    }, methods: {
        getSaml(id) {
            axios.get(`/auth-manager/get-saml-response/${id}`).then(res => {
                const samlRes = JSON.parse(res.data)
                this.saml = samlRes.Response
            }).catch(err =>
                console.log(err.response)
            )
        }
    }
}
</script>
