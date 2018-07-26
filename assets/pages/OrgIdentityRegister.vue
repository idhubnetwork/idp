<template>
    <div>
        <div class="tile is-ancestor">
            <div class="tile is-parent">
                <div class="tile is-child box">
                    <p class="title">{{ $t('org-identity-register') }}</p>

                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Readable Name or Role Name</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                                <div class="control">
                                    <input class="input" type="text" v-model="roles" placeholder="e.g. I'm so rich...">
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Provider ARN</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                                <div class="control">
                                    <input class="input" type="text" v-model="providerArn" placeholder="e.g. arn:aws:iam::000000000000:saml-provider/IDHub">
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Role ARN</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                                <div class="control">
                                    <input class="input" type="text" v-model="roleArn" placeholder="e.g. arn:aws:iam::000000000000:role/IDHUB_IDP">
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="field is-horizontal">
                        <div class="field-label is-normal">
                            <label class="label">Org. AWS Identity</label>
                        </div>
                        <div class="field-body">
                            <div class="field">
                                <div class="control">
                                    <input class="input" type="text" v-model="awsId" placeholder="e.g. 000000000000">
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="field is-horizontal">
                        <div class="field-label"></div>
                        <div class="field-body">
                            <div class="field">
                                <div class="control">
                                    <button class="button is-link is-rounded" @click="register">{{ $t('register') }}</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { axios } from '@/util'

export default {
    data() {
        return {
            roles: '',
            providerArn: '',
            roleArn: '',
            awsId: ''
        }
    }, methods: {
        register() {
            axios.post('/org-user/create', {
                roles: this.roles,
                providerArn: this.providerArn,
                roleArn: this.roleArn,
                awsId: this.awsId
            }).then(res =>
                console.log(res.data)
            ).catch(err =>
                console.log(err.response.data)
            )
        }
    }
}
</script>
