// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import Vue from 'vue'
import axios from 'axios'
import store from '@/store'
import VueI18n from 'vue-i18n'

const API_URL = process.env.API_URL

Vue.use(VueI18n)

const i18n = new VueI18n({
  locale: 'en',
  messages: {}
})

/**
 * @param {String} locale
 */
export async function loadMessages (locale) {
  if (Object.keys(i18n.getLocaleMessage(locale)).length === 0) {
    // const messages = await import(/* webpackChunkName: "lang-[request]" */ `@/lang/${locale}`)
    const { data } = await axios.get(API_URL + `/lang/${locale}.json`)
    i18n.setLocaleMessage(locale, data)
  }

  if (i18n.locale !== locale) {
    i18n.locale = locale
  }
}

(async function () {
  await loadMessages(store.getters['lang/locale'])
})()

export default i18n
