<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="admin-main" v-if="!$route.query.preview">
    <Sidebar />
    <Share ref="share" v-if="$route.query.share" />
    <Newpage ref="newpage" v-if="user && user.access_level > 0 && $route.query.new" />
    <Duplicate ref="duplicate" v-if="user && user.access_level > 0 && $route.query.duplicate" />
    <Version ref="version" v-if="user && user.access_level > 0 && $route.query.version" />
    <transition name="fade">
      <Settings ref="settings" v-if="user && user.access_level > 0 && $route.query.settings" />
    </transition>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import swal from 'sweetalert2'
import axios from 'axios'
import Sidebar from './Sidebar.vue'
import Share from './Share.vue'
import Newpage from './Newpage.vue'
import Duplicate from './Duplicate.vue'
import Version from './Version.vue'
import Settings from './Settings.vue'
export default {
  name: 'admin-main',

  components: {
    Sidebar,
    Newpage,
    Version,
    Duplicate,
    Share,
    Settings
  },

  computed: mapGetters({
    pages: 'page/current_pages',
    user: 'auth/user',
    check: 'auth/check'
  }),

  data: () => ({
    whitelist: ['share']
  }),

  watch: {
    // Watches address query.
    async '$route' (v) {
      if (Object.keys(v.query).length > 0 && this.whitelist.indexOf(Object.keys(v.query)[0]) === -1) {
        if (!this.$store.getters['auth/check'] && this.$store.getters['auth/token']) {
          try {
            await this.$store.dispatch('auth/fetchUser')
          } catch (e) {
            this.$router.replace({name: 'Login', query: {redirect: this.$route.fullPath}})
          }
        } else if (!this.$store.getters['auth/token']) {
          this.$router.replace({name: 'Login', query: {redirect: this.$route.fullPath}})
        }
      }
      this.$nextTick(() => {
        if (v.query.archive) {
          this.archivePageInit()
        }
        if (v.query.publish) {
          this.publishPageInit()
        }
        if (v.query.logout) {
          this.logout()
        }
      })
    },
    'pages.details' () {
      if (this.$route.query.archive) {
        this.archivePageInit()
      }
      if (this.$route.query.publish) {
        this.publishPageInit()
      }
    }
  },

  methods: {
    // Confirms page archiving.
    async archivePageInit () {
      if (this.pages.details.id && this.check) {
        swal({
          title: this.$t('settings_pages__archive_warning_1'),
          text: this.$t('settings_pages__archive_warning_2') + this.pages.details.name + this.$t('settings_pages__archive_warning_3'),
          type: 'warning',
          showCancelButton: true,
          confirmButtonColor: '#343a40',
          cancelButtonColor: '#6c757d',
          confirmButtonText: this.$t('button__yes_archive'),
          cancelButtonText: this.$t('button__no')
        }).then((result) => {
          if (result.value) {
            this.archivePage()
          } else {
            this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
          }
        })
      }
    },

    // Archives page.
    async archivePage () {
      try {
        await axios.put(process.env.API_URL + '/api/page/' + this.pages.details.id + '/archive', {
          headers: {Authorization: `Bearer ${this.token}`}
        })
        this.$awn.success(this.$t('settings_pages__page_archived'))
        this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
        this.$root.$refs['app-layout'].initData() // Updates menu icons.
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },

    // Confirms page publishing.
    async publishPageInit () {
      if (this.pages.details.id && this.check) {
        swal({
          title: this.$t('settings_pages__publish_page_warning_1'),
          text: this.$t('settings_pages__publish_page_warning_2') + this.pages.details.name + this.$t('settings_pages__publish_page_warning_3'),
          type: 'warning',
          showCancelButton: true,
          cancelButtonColor: '#6c757d',
          confirmButtonColor: '#ffc107',
          cancelButtonText: this.$t('button__no'),
          confirmButtonText: this.$t('button__yes_publish')
        }).then((result) => {
          if (result.value) {
            this.publishPage()
          } else {
            this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
          }
        })
      }
    },

    // Publishes page.
    async publishPage () {
      try {
        await axios.put(process.env.API_URL + '/api/page/' + this.pages.details.id + '/publish', {
          headers: {Authorization: `Bearer ${this.token}`}
        })
        this.$awn.success(this.$t('settings_pages__page_published'))
        this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
        this.$root.$refs['app-layout'].initData() // Updates menu icons.
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },

    // Logs out the user.
    logout () {
      this.$store.dispatch('auth/logout')
      this.$router.replace('/')
      this.$root.$refs['app-layout'].initData()
    }
  }
}
</script>
