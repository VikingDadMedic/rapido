// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Félix PORTIER <f.portierdev@protonmail.com>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import axios from 'axios'
import Cookies from 'js-cookie'

import * as types from '../mutation-types.js'
const API_URL = process.env.API_URL


// State
export const state = {
  user: null,
  users: [],
  user_level: [],
  token: Cookies.get('token')
}

// Metters
export const getters = {
  user: state => state.user,
  users: state => state.users,
  user_level: state => state.user_level,
  token: state => state.token,
  check: state => state.user !== null
}

// Mutations
export const mutations = {
  [types.SAVE_TOKEN] (state, { token, remember }) {
    state.token = token
    Cookies.set('token', token, { expires: remember ? 365 : null })
  },

  [types.FETCH_USER_SUCCESS] (state, { user }) {
    state.user = user
  },

  [types.FETCH_ALL_USER] (state, payload) {
    state.users = payload
  },

  [types.FETCH_USER_LEVEL] (state, payload) {
    state.user_level = payload
  },

  [types.FETCH_USER_FAILURE] (state) {
    state.token = null
    Cookies.remove('token')
  },

  [types.LOGOUT] (state) {
    state.user = null
    state.token = null

    Cookies.remove('token')
  },

  [types.UPDATE_USER] (state, { user }) {
    state.user = user
  }
}

// actions
export const actions = {
  saveToken ({ commit, dispatch }, payload) {
    commit(types.SAVE_TOKEN, payload)
  },

  async fetchUser ({ state, commit }) {
    try {
      const { data } = await axios.get(API_URL + '/api/check_token', {headers: {Authorization: `Bearer ${state.token}`}})

      commit(types.FETCH_USER_SUCCESS, { user: data })
    } catch (e) {
      commit(types.FETCH_USER_FAILURE)
    }
  },

  async fetchAllUser ({ commit, dispatch }) {
    try {
      const { data } = await axios.get(API_URL + '/api/account', {headers: {Authorization: `Bearer ${state.token}`}})

      dispatch('setUserLevel')
      commit(types.FETCH_ALL_USER, data)
    } catch (e) {
      commit(types.FETCH_ALL_USER, [])
    }
  },

  setUserLevel ({ commit }) {
    var obj = [
      {value: 1, name: 'settings_accounts__user'},
      {value: 10, name: 'settings_accounts__administrator'}
    ]
    commit(types.FETCH_USER_LEVEL, obj)
  },

  updateUser ({ commit }, payload) {
    commit(types.UPDATE_USER, payload)
  },

  logout ({ commit }, payload) {
    commit(types.LOGOUT)
  },

  async fetchOauthUrl (ctx, { provider }) {
    const { data } = await axios.post(`/api/oauth/${provider}`)

    return data.url
  },

  async fetchSetting ({state}) {
    const { data } = await axios.get(API_URL + '/api/setting', {headers: {Authorization: `Bearer ${state.token}`}})

    return data
  }
}
