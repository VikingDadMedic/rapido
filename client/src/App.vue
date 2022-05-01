<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div id="app">
    <loading ref="loading"></loading>
    <transition name="page" mode="out-in">
      <component v-if="layout" :is="layout" :ref="layout.name"></component>
    </transition>
    <gallery-modal />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Loading from '@/components/Loading.vue'

// Load layout components dynamically.
const requireContext = require.context('@/layouts', false, /.*\.vue$/)

const layouts = requireContext.keys()
  .map(file => [file.replace(/(^.\/)|(\.vue$)/g, ''), requireContext(file)])
  .reduce((components, [name, component]) => {
    components[name] = component
    return components
  }, {})

const appName = process.env.appName

export default {
  el: '#app',

  components: {
    Loading
  },

  metaInfo () {
    return {
      title: '',
      titleTemplate: `%s · ${this.logo.name || appName}`
    }
  },

  computed: {
    ...mapGetters({
      logo: 'page/logo'
    })
  },

  data: () => ({
    layout: null,
    defaultLayout: 'app',
    settings: {}
  }),

  async beforeMount () {
    await this.$store.dispatch('lang/getLocales')
  },

  mounted () {
    this.$loading = this.$refs.loading
  },

  methods: {
    /**
     * Set the application layout.
     *
     * @param {String} layout
     */
    setLayout (layout) {
      if (!layout || !layouts[layout]) {
        layout = this.defaultLayout
      }

      this.layout = layouts[layout].default
    }
  }
}
</script>

<style lang="stylus" src="@/assets/stylus/app.styl"></style>
