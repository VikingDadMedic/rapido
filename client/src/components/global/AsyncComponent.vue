<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="componentWrapper">
    <component v-if="componentLoader" :is="componentLoader" :api_url="api_url" :manifest="manifest">
      <!-- {{ this.$props.slot }} -->
    </component>
  </div>
</template>

<script>
import axios from 'axios'
import { mapGetters } from 'vuex'
export default {
  name: 'async-component',
  props: {
    componentFile: {
      type: String,
      default: () => null
    },
    options: {
      type: Array,
      default: () => []
    },
    extension: {
      type: String,
      required: true
    },
    api_url: {
      type: String,
      default: ''
    }
  },
  mounted () {
    this.getManifest()
  },
  computed: {
    ...mapGetters({
      token: 'auth/token'
    }),
    componentLoader () {
      return window.ContactForm ? window[this.$toTitleCase(this.$toCamelCase(this.componentFile))].default : null
    }
  },
  data: () => ({
    manifest: {}
  }),
  methods: {
    async getManifest () {
      let file = this.extension.split('/')
      const { data } = await axios.get(process.env.API_URL + '/extension/' + this.extension + '/' + file[2] + '.manifest.json')

      this.$set(this, 'manifest', data)
    }
  }
}
</script>
