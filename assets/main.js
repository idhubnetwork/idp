'use strict'

import 'bulma'
import 'font-awesome'
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueI18n from 'vue-i18n'

import App from '@/App.vue'

import Index from '@/pages/Index.vue'
import Signin from '@/pages/Signin.vue'
import MetadataManager from '@/pages/MetadataManager.vue'
import OrgIdentity from '@/pages/OrgIdentity.vue'
import OrgIdentityRegister from '@/pages/OrgIdentityRegister.vue'
import OrgIdentityPendingList from '@/pages/OrgIdentityPendingList.vue'
import AuthManager from '@/pages/AuthManager.vue'
import AuthManagerList from '@/pages/AuthManagerList.vue'
import AuthManagerApply from '@/pages/AuthManagerApply.vue'

Vue.use(VueRouter)
Vue.use(VueI18n)

const router = new VueRouter({
    linkActiveClass: 'is-active',
    routes: [
        { path: '/', component: Index },
        { path: '/signin', component: Signin },
        { path: '/metadata-manager', component: MetadataManager },
        {
            path: '/org-identity', component: OrgIdentity, children: [
                { path: 'register', component: OrgIdentityRegister },
                { path: 'pending-list', component: OrgIdentityPendingList }
            ]
        },
        {
            path: '/auth-manager', component: AuthManager, children: [
                { path: 'list', component: AuthManagerList },
                { path: 'apply', component: AuthManagerApply }
            ]
        }
    ],
})

router.beforeEach((to, from, next) => {
    if (to.path === '/' || to.path === '/signin' || (document.cookie.indexOf('IDHUB_IDENTITY=') >= 0)) {
        next()
    } else {
        next('/signin')
        location.reload()
    }
})

const i18n = new VueI18n({
    locale: 'zh_CN',
    fallbackLocale: 'zh_CN',
    messages: {
        'zh_TW': require('@/i18n/zh_TW.json'),
        'zh_CN': require('@/i18n/zh_CN.json'),
        'en_US': require('@/i18n/en_US.json')
    }
})

new Vue({
    router,
    i18n,
    template: '<App></App>',
    components: { App }
}).$mount('#app')