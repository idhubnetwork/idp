<template>
    <div>
        <h1 class="title">{{ $t('org-identity') }}</h1>

        <div class="tabs is-centered is-boxed">
            <ul>
                <li>
                    <router-link class="navbar-item" to="/org-identity/register" exact>{{ $t('register') }}</router-link>
                </li>
            </ul>
            <ul>
                <li>
                    <router-link class="navbar-item" to="/org-identity/pending-list" exact>{{ $t('pending-list') }}</router-link>
                </li>
            </ul>
        </div>

        <router-view></router-view>
    </div>
</template>

<style lang="scss" scoped>
a.is-active {
  border-bottom-color: #3273dc;
  color: #3273dc;
}
</style>