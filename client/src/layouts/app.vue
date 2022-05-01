<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="app-layout">
    <admin />
    <div class="template-wrapper">
      <child />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import _ from 'lodash'
import Admin from '@/components/admin/Index.vue'

export default {
  name: 'app-layout',

  // Set title meta information
  metaInfo () {
    let obj = {
      title: this.$route.params.page || 'home'.replace(/\b\w/g, l => l.toUpperCase())
    }
    if (this.logo.favicon_url) {
      obj.link = [
        { rel: 'icon', href: this.logo.favicon_url || '' }
      ]
    }
    return obj
  },

  computed: {
    ...mapGetters({
      logo: 'page/logo'
    })
  },

  mounted () {
    this.initData()
  },

  watch: {
    '$route.path' (v) {
      this.initData()
    }
  },

  components: {
    Admin
  },

  methods: {
    // fetch initialization page data
    initData: _.debounce(function () {
      this.$nextTick(async () => {
        const data = await this.$store.dispatch('page/init', {page: this.$route.path.slice(1) || ''})
        if (!data) {
          this.$router.replace('/404')
        }
      })
    }, 100)
  }
}
</script>
