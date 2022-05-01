<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>,Félix PORTIER <f.portierdev@prontonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->



<template>
  <li @mouseover="dropDownShow(true)" @mouseleave="dropDownShow(false)" class="nav-item dropdown">
    <a v-if="!currLocale" id="current-locale" class="nav-link dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ currentLocale }}
    </a>
    <a v-else id="current-locale" class="nav-link dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ currLocale }}
    </a>

    <!--ADAPT TO LENGTH OF lang.json ( currently 40 values ) won't display properly if lenght different than 40-->
    <!--TO DO ? MAKE DROPDOWN ADAPT AUTOMATICALLY ? -->
    <template>
      <div id="locales-lang" class="dropdown-menu dropdown-multicol lang" style="left: -52em; top: 0em  !important">
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(0,5) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(5,10)" :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(10,15) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(15,20) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(20,25) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(25,30) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
        <div class="dropdown-col">
          <a v-for="locale in locales.slice(35,40) " :key="locale.code" class="dropdown-item locale" @click.prevent="setLocale(locale.code,locale.name)">
            {{ locale.name }}
          </a>
        </div>
      </div>
    </template>

  </li>
</template>

<script>
import Cookies from 'js-cookie'
import { mapGetters } from 'vuex'
import { loadMessages } from '@/plugins/i18n.js'
import {getSupportedLocales} from '@/store/modules/supported-locales.js'

// import _ from 'lodash'

export default {
  computed: {
    ...mapGetters({currentLocale: 'lang/localeName'})
  },

  data: () => ({
    // Gets supportedLocales values into annotated array.
    locales: getSupportedLocales(),
    currLocale: this.currentLocale
  }),

  mounted () {
    this.currentFill()
  },

  updated () {
    this.currentFill()
  },
  methods: {
    currentFill () {
      const elems = document.getElementsByClassName('dropdown-item locale')
      for (let i = 1; i < elems.length; i++) {
        if (elems[i].innerHTML.replace(/\s+/g, '') === document.querySelector('#current-locale').innerHTML.replace(/\s+/g, '')) {
          elems[i].classList.add('hovered')
        }
      }
    },
    dropDownShow (bool) {
      if (bool) {
        document.getElementById('locales-lang').classList.add('show')
      } else {
        document.getElementById('locales-lang').classList.remove('show')
      }
    },

    // setLocale(locale,val) sets i18n locale to the one selected in the dropdown.
    setLocale (locale, val) {
      if (this.$i18n.locale !== locale) {
        loadMessages(locale)
        this.$store.dispatch('lang/setLocale', {locale})
        Cookies.set('localeName', val, {expires: 365, sameSite: 'strict'})
        const elems = document.getElementsByClassName('dropdown-item locale')
        for (let i = 0; i < elems.length; i++) {
          if (elems[i].innerHTML.replace(/\s+/g, '') === document.querySelector('#current-locale').innerHTML.replace(/\s+/g, '')) {
            elems[i].classList.remove('hovered')
          }
          if (elems[i].innerHTML.replace(/\s+/g, '') === val) {
            elems[i].classList.add('hovered')
          } else {
            elems[i].classList.remove('hovered')
          }
        }
      }
      this.currLocale = val
      this.dropDownShow(false)
    }
  }
}
</script>

<style lang="stylus" scoped>

.dropdown-menu
    margin-top: 0
    width: auto
  // .dropdown-toggle: active
  // pointer-events: none

.dropdown-multicol
  > .dropdown-col
    display : inline-block

@media (max-width: 990px) {
  .show {
    flex-direction: column
  }
}
.show
  display : flex

.dropdown-item
  color: rgba(0,0,0,.5)
.dropdown-item:hover
  background: #28a745
  color: #ffffff
.dropdown-item:focus
  background: unset
.hovered
  color: #28a745

</style>
