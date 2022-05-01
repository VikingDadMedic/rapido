<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <nav v-bind:style="{display: navVisibility}" ref="navbar" class="navbar fixed-top navbar-fixed-top navbar-expand-lg" :class="manifest.manifest ? 'navbar-default' : 'navbar-light'">
    <router-link v-bind:style="{display: logoVisibility}" to="/" class="navbar-brand">
      <img class="navbar-logo" :src="logo.logo_url" style="max-height: 30px;" v-if="logo.logo_url">
      <template v-else>
        {{ appName }}
      </template>
    </router-link>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <MenuList v-bind:style="{display: menusVisibility}" id="admin-menu-header" :navs="menus.header" location="header" />
      <ul class="navbar-nav ml-auto">
        <locale-dropdown v-bind:style="{display: localeDropdownVisibility}" />
      </ul>
    </div>
  </nav>
</template>

<script>

import { mapGetters } from 'vuex'
import LocaleDropdown from './LocaleDropdown.vue'
import _ from 'lodash'

const appName = process.env.appName

export default {
  name: 'navbar',

  components: {
    LocaleDropdown
  },

  data: () => ({
    appName: appName,
    items: {
      show_logo: false,
      show_menus: false,
      show_locales_menu: false
    },
    navVisibility: 'none',
    logoVisibility: 'none',
    menusVisibility: 'none',
    localeDropdownVisibility: 'none'
  }),

  computed: {
    ...mapGetters({
      manifest: 'page/manifest',
      menus: 'page/menus',
      logo: 'page/logo'
    })
  },

  mounted () {
    this.getNavStyles()
  },
  updated () {
    this.getNavStyles()
  },
  methods: {
    async getNavStyles () {
      try {
        const data = await this.$store.dispatch('auth/fetchSetting')
        _.assign(this.items, data)
        if (data.show_logo || data.show_menus || data.show_locales_menu) {
          this.navVisibility = 'flex'
        }
        if (data.show_logo) {
          this.logoVisibility = 'flex'
        }
        if (data.show_menus) {
          this.menusVisibility = 'flex'
        }
        if (data.show_locales_menu) {
          this.localeDropdownVisibility = 'flex'
        }
      } catch (e) {
      }
    }
  }
}
</script>
