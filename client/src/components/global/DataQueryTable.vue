<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <tbody :id="parent">
    <tr>
      <td :colspan="columnspan">
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text bg-transparent">
              <span class="fa fa-search"></span>
            </span>
          </div>
          <input type="text" class="form-control form-control-lg" :placeholder="$t('search_engine__search')" v-model="query.key">
        </div>
      </td>
    </tr>
    <template v-for="(dt, key, idx) in datas" v-if="datas.length > 0">
      <tr :key="idx">
        <slot name="items" :data="dt" :keys="key" :index="idx">
          <td>
            {{ dt }}
          </td>
        </slot>
      </tr>
    </template>
    <tr v-if="datas.length === 0">
      <td :colspan="columnspan">
        <h4 class="text-center">
          {{ $t('search_engine__no_match') }}
        </h4>
      </td>
    </tr>
    <tr>
      <td :colspan="columnspan">
        <ul class="pagination">
          <li class="page-item" v-if="this.query.current !== 1">
            <a class="page-link" aria-label="Previous" @click="prev">
              <span aria-hidden="true">
                &laquo;
              </span>
              <span class="sr-only">
                {{ $t('button__previous_page') }}
              </span>
            </a>
          </li>
          <li class="page-item" v-for="(i, key) in pages" :key="key" :class="{'active': i === query.current}">
            <a class="page-link" @click="changeCurrent(i)">
              {{ i }}
            </a>
          </li>
          <li class="page-item" v-if="this.query.current !== parseInt(this.pages[this.pages.length - 1]) && datas.length > 0">
            <a class="page-link" aria-label="Next" @click="next">
              <span aria-hidden="true">
                &raquo;
              </span>
              <span class="sr-only">
                {{ $t('button__next_page') }}
              </span>
            </a>
          </li>
        </ul>
      </td>
    </tr>
  </tbody>
</template>

<script>
import _ from 'lodash'
export default {
  name: 'DataQueryTable',

  props: {
    data: {
      type: Array,
      required: true
    },
    query: {
      type: Object,
      default: () => ({
        key: null,
        column: null,
        sort: null,
        sortColumn: null,
        limit: 10,
        offset: 0,
        current: 1
      })
    },
    parent: {
      type: String,
      default: null
    },
    columnspan: {
      type: Number,
      default: 0
    }
  },

  mounted () {
    this.datas = true
  },

  computed: {
    datas: {
      get () {
        return this.setData()
      },
      set () {
        this.setPages()
      }
    }
  },

  watch: {
    data () {
      this.datas = true
      this.changeCurrent(1)
    },
    datas () {
      this.datas = true
      // this.changeCurrent(1)
    },
    'query.current' (v) {
      this.$set(this.query, 'offset', (v - 1) * this.query.limit)
    },
    'query.key' (v) {
      this.changeCurrent(1)
    }
  },

  data: () => ({
    totalData: 0,
    pages: []
  }),

  methods: {
    setData () {
      let result = []
      let data = JSON.parse(JSON.stringify(this.data))
      if (this.query.column && this.query.key) {
        data = _.filter(data, (o) => {
          let val = this.query.column.split(',')
          let res = []
          for (let i in val) {
            if (o[val[i]]) {
              res.push(o[val[i]].toLowerCase().includes(this.query.key.toLowerCase()))
            }
          }
          if (res.indexOf(true) !== -1) {
            return true
          } else {
            return false
          }
        })
      }
      for (let i = this.query.offset; i < (this.query.limit * this.query.current); i++) {
        if (data[i]) {
          result.push(data[i])
        }
      }
      this.$set(this, 'totalData', data.length)
      return result
    },
    setPages () {
      let numbers = []
      let total = Math.ceil(this.totalData / this.query.limit)
      for (let i = 0; i < total; i++) {
        numbers.push(i + 1)
      }
      this.$set(this, 'pages', numbers)
    },
    changeCurrent (v) {
      this.$set(this.query, 'current', v)
    },
    next () {
      this.changeCurrent(this.query.current + 1)
    },
    prev () {
      this.changeCurrent(this.query.current - 1)
    }
  }
}
</script>

<style lang="stylus" scoped>
.input-group
  .input-group-text, .form-control
    border: 0;
    border-bottom: 1px solid grey
</style>
