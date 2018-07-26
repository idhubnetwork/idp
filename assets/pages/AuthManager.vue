<template>
    <div>
        <h1 class="title">{{ $t('auth-manager') }}</h1>

        <div class="tabs is-centered is-boxed">
            <ul>
                <li>
                    <router-link class="navbar-item" to="/auth-manager/list" exact>{{ $t('auth-manager-list') }}</router-link>
                </li>
            </ul>
            <ul>
                <li>
                    <router-link class="navbar-item" to="/auth-manager/apply" exact>{{ $t('auth-manager-apply') }}</router-link>
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
