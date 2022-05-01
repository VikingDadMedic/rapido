// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>,Felix PORTIER <fportierdev@protonmail.com>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


// import axios from 'axios'
import Cookies from 'js-cookie'
import * as types from '../mutation-types.js'

// const API_URL = process.env.API_URL

const { locale } = {
  locale: 'eng'
}

// state
export const state = {
  localeName: Cookies.get('localeName'),
  locale: Cookies.get('locale') || locale,
  locales: {}
}

// getters
export const getters = {
  locale: state => state.locale,
  locales: state => state.locales,
  localeName: state => state.localeName

}

// mutations
export const mutations = {
  [types.SET_LOCALE] (state, { locale }) {
    state.locale = locale
  },
  [types.SET_LOCALES] (state, locales) {
    state.locales = locales
  }
}

// actions
export const actions = {
  setLocale ({ commit }, { locale }) {
    commit(types.SET_LOCALE, { locale })
    Cookies.set('locale', locale, {expires: 365, sameSite: 'strict'})
  }
}
