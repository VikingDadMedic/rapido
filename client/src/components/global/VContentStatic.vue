<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div :class="col" v-for="(pg, key) in content" :key="key" v-if="content.length > 0">
      <div class="ql-editor page-content">
        <template v-if="!pg.extension">
          <div v-html="pg.content.content"></div>
        </template>
        <template v-else>
          <ContactForm :to="pg.json_settings[0].value" />
        </template>
      </div>
    </div>
  </div>
</template>

<script>
import AsyncComponent from './AsyncComponent.vue'
export default {
  name: 'VContentStatic',

  metaInfo () {
    let obj = []
    for (let i in this.content) {
      if (this.content[i].extension) {
        let file = this.content[i].extension.split('/')
        obj.push({ src: process.env.API_URL + '/extension/' + this.content[i].extension + '/dist/umd/' + file[2] + '.min.js', type: 'application/javascript' })
      }
    }
    return {
      script: obj
    }
  },

  components: {
    AsyncComponent
  },

  props: {
    content: {
      type: Array,
      default: () => ([])
    },
    col: {
      type: String,
      default: 'col-md-12'
    }
  },

  data: () => ({
    isLoaded: {},
    api_url: process.env.API_URL
  }),

  watch: {
    content (v) {
      this.initCheckComponent(v)
    }
  },

  mounted () {
    this.initCheckComponent(this.content)
  },

  methods: {
    initCheckComponent (v) {
      for (let i in v) {
        if (v[i].extension) {
          // this.checkComponent(v[i], i)
        }
      }
    }
    /* checkComponent (v, idx) {
      if (v.extension) {
        let check = setInterval(() => {
          if (window[this.$toTitleCase()]) {
            this.$set(this.isLoaded, idx, true)
            clearInterval(check)
          }
        }, 500)
      }
    } */
  }
}
</script>
