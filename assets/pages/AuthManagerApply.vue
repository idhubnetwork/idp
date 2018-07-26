<template>
    <div>
        <p class="title">{{ $t('auth-manager-apply') }}</p>

        <table class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth">
            <thead>
                <tr>
                    <th>#</th>
                    <th>identity</th>
                    <th>roles</th>
                    <th>created at</th>
                    <th>updated at</th>
                    <th>action</th>
                    <th>apply status</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(orgUser, index) in orgUsers" :key="orgUser.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ orgUser.identity }}</td>
                    <td>{{ orgUser.roles }}</td>
                    <td>{{ orgUser.createdAt }}</td>
                    <td>{{ orgUser.updatedAt }}</td>
                    <td>
                        <button class="button is-link is-rounded is-small" v-if="orgUser.userIdentity === null" @click="apply(orgUser)">apply</button>
                    </td>
                    <td>{{ orgUser.isActive }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
import { axios } from '@/util'

export default {
    data() {
        return {
            orgUsers: []
        }
    }, mounted() {
        axios.get(`/auth-manager/apply-list`).then(res => {
            this.orgUsers = res.data
        }).catch(err =>
            console.log(err.response)
        )
    }, methods: {
        apply(orgUser) {
            axios.post(`/auth-manager/apply`, {
                orgUser: orgUser.identity,
                roles: orgUser.roles,
            }).then(res => {
                location.reload()
            }).catch(err =>
                console.log(err.response)
            )
        }
    }
}
</script>