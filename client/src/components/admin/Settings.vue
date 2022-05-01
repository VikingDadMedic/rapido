<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="admin-settings">
    <div class="admin-settings-container">
      <div class="admin-settings-container-inner h-100">
        <div class="template-wrapper h-100">
          <div class="template-header">
            <div class="container">
              <nav class="navbar fixed-top navbar-fixed-top navbar-expand-lg template-header-menu" :class="manifest.manifest ? 'navbar-default' : 'navbar-light'">
                <div class="text-center">
                  <a class="navbar-brand" style="cursor: default;">
                    <span class="fa fa-wrench" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('sidebar__settings') }}
                  </a>
                </div>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarAdmin" aria-controls="navbarAdmin" aria-expanded="false" aria-label="Toggle navigation">
                  <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarAdmin">
                  <ul class="navbar-nav mr-auto">
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'pages'}}" class="nav-item" exact>
                      <a class="nav-link">
                        <span class="fa fa-file" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_pages') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'contents'}}" class="nav-item" exact>
                      <a class="nav-link">
                        <span class="fa fa-th-large" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_contents') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'files'}}" class="nav-item" exact>
                      <a class="nav-link">
                        <span class="fa fa-image" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_files') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'features'}}" class="nav-item" exact >
                      <a class="nav-link">
                        <span class="fa fa-puzzle-piece" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_features') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'accounts'}}" class="nav-item" exact>
                      <a class="nav-link">
                        <span class="fa fa-user" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_accounts') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'system'}}" class="nav-item" exact v-if="user.access_level >= 10">
                      <a class="nav-link">
                        <span class="fa fa-cogs" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_system') }}</span>
                      </a>
                    </router-link>
                    <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {'settings': 'about'}}" class="nav-item" exact>
                      <a class="nav-link">
                        <span class="fa fa-info" aria-hidden="true"></span>&nbsp;&nbsp;<span class="menu-title">{{ $t('settings_about') }}</span>
                      </a>
                    </router-link>
                  </ul>
                  <ul class="navbar-nav ml-lg-auto d-flex align-items-center flex-row">
                    <li class="nav-item" @click="closePanel">
                      <a class="nav-link" aria-hidden="true">
                        <span class="fa fa-times"></span>&nbsp;&nbsp;<span class="menu-title" aria-hidden="true">{{ $t('button__close') }}</span>
                      </a>
                    </li>
                  </ul>
                </div>
              </nav>
            </div>
          </div>

          <!-- partial -->
          <div class="template-main">
            <div class="container py-5">
              <div class="row justify-content-center">
                <Child>
                  <div class="col-12">
                    <transition name="page" mode="out-in">
                      <keep-alive>
                        <component v-if="content" :is="content"></component>
                      </keep-alive>
                    </transition>
                  </div>
                </Child>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
import $ from 'jquery'
import { mapGetters } from 'vuex'

// Load content components dynamically.
const requireContext = require.context('./Settings', false, /.*\.vue$/)

const contents = requireContext.keys()
  .map(file => [file.replace(/(^.\/)|(\.vue$)/g, ''), requireContext(file)])
  .reduce((components, [name, component]) => {
    components[name.toLowerCase()] = component
    return components
  }, {})

export default {
  name: 'admin-settings',

  computed: {
    ...mapGetters({
      manifest: 'page/manifest',
      user: 'auth/user'
    })
  },

  mounted () {
    this.$nextTick(() => {
      if (this.$route.query.settings) {
        this.setBg()
        $('body').css('overflow', 'hidden')
        this.setContent(this.$route.query.settings)
      }
    })
  },

  data: () => ({
    content: null,
    defaultContent: 'pages'
  }),

  beforeDestroy () {
    $('body').attr('style', '')
  },

  watch: {
    '$route.query.settings' (v) {
      if (v) {
        this.setBg()
        this.setContent(v)
      }
    }
  },

  methods: {
    async initData () {
      await this.$store.dispatch('page/initSettings', {})
    },
    setBg () {
      this.$nextTick(() => {
        setTimeout(() => {
          this.$el.querySelector('.admin-settings-container-inner').style.backgroundColor = $('body').css('background-color')
        }, 500)
      })
    },
    /**
     * Set the application content.
     *
     * @param {String} content
     */
    setContent (content) {
      if (!content || !contents[content]) {
        this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}, query: {settings: this.defaultContent}})
      }

      this.content = contents[content].default
    },
    sidebarToggle () {
      $('.admin-settings-container').toggleClass('sidebar-icon-only')
    },
    closePanel () {
      this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
    }
  }
}
</script>
