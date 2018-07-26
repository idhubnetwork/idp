<template>
    <div>
        <p class="title">{{ $t('pending-list') }}</p>

        <table class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth">
            <thead>
                <tr>
                    <th>#</th>
                    <th>roles</th>
                    <th>identity</th>
                    <th>created at</th>
                    <th>updated at</th>
                    <th colspan="2">actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(user, index) in users" :key="user.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ user.roles }}</td>
                    <td>{{ user.identity }}</td>
                    <td>{{ user.createdAt }}</td>
                    <td>{{ user.updatedAt }}</td>
                    <td>
                        <button class="button is-success is-rounded is-small" @click="approval(user.id, 1)">approved</button>
                    </td>
                    <td>
                        <button class="button is-danger is-rounded is-small" @click="approval(user.id, -1)">rejected</button>
                    </td>
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
            users: []
        }
    }, mounted() {
        axios.get(`/org-user/pending-list`).then(res => {
            this.users = res.data
        }).catch(err =>
            console.log(err.response)
        )
    }, methods: {
        pendingList() {

        }, approval(id, isActive) {
            axios.post(`/org-user/approval`, {
                id: id,
                isActive: isActive
            }).then(res => {
                const index = this.users.findIndex(e => e.id === id)
                this.users.splice(index, 1)
            }).catch(err =>
                console.log(err.response)
            )
        }
    }
}
</script>
