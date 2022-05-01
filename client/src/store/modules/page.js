// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import axios from 'axios'
import store from '@/store'
import * as types from '../mutation-types.js'

const API_URL = process.env.API_URL

// state
export const state = {
  access_level: [],
  logo: {
    favicon_url: null,
    logo_url: null,
    name: null
  },
  published_pages: [],
  archived_pages: [],
  contents: [],
  contactforms: [],
  current_pages: {
    details: {},
    content: [],
    version: []
  },
  preview: localStorage.getItem('preview') ? JSON.parse(localStorage.getItem('preview')) : {
    details: {},
    content: []
  },
  menus: {
    header: {},
    main: {},
    footer: {}
  },
  themes: [],
  system: [],
  manifest: {}
}

// getters
export const getters = {
  access_level: state => state.access_level,
  logo: state => state.logo,
  published_pages: state => state.published_pages,
  archived_pages: state => state.archived_pages,
  contents: state => state.contents,
  current_pages: state => state.current_pages,
  preview: state => state.preview,
  menus: state => state.menus,
  contactforms: state => state.contactforms,
  themes: state => state.themes,
  system: state => state.system,
  manifest: state => state.manifest
}

// mutations
export const mutations = {
  [types.FETCH_ACCESS_LEVEL] (state, payload) {
    state.access_level = payload
  },
  [types.FETCH_LOGO] (state, payload) {
    state.logo = payload
  },
  [types.FETCH_PUBLISHED_PAGES] (state, payload) {
    state.published_pages = payload
  },
  [types.FETCH_ARCHIVED_PAGES] (state, payload) {
    state.archived_pages = payload
  },
  [types.FETCH_ALL_CONTENTS] (state, payload) {
    state.contents = payload
  },
  [types.FETCH_ALL_CONTACTFORMS] (state, payload) {
    state.contactforms = payload
  },
  [types.FETCH_CURRENT_PAGES_CONTENTS] (state, payload) {
    state.current_pages.content = payload
  },
  [types.FETCH_CURRENT_PAGES_DETAILS] (state, payload) {
    state.current_pages.details = payload
  },
  [types.FETCH_CURRENT_PAGES_VERSIONS] (state, payload) {
    state.current_pages.version = payload
  },
  [types.FETCH_ALL_MENUS] (state, payload) {
    if (payload) {
      for (var i in payload) {
        state.menus[payload[i].name] = payload[i]
      }
    } else {
      state.menus.header = {
        item: []
      }
      state.menus.main = {
        item: []
      }
      state.menus.footer = {
        item: []
      }
    }
  },
  [types.FETCH_PREVIEW_CONTENT] (state, payload) {
    localStorage.removeItem('preview')
    state.preview = payload
    localStorage.setItem('preview', JSON.stringify(payload))
  },
  [types.FETCH_ALL_THEMES] (state, payload) {
    state.themes = payload
  },
  [types.FETCH_ALL_SYSTEM] (state, payload) {
    state.system = payload
  },
  [types.FETCH_MANIFEST] (state, payload) {
    state.manifest = payload
  }
}

// actions
export const actions = {
  async init ({ dispatch }, payload) {
    dispatch('setAccessLevel')
    dispatch('fetchPublishedPages')
    dispatch('fetchArchivedPages')
    dispatch('fetchAllContent')
    dispatch('fetchAllContactForm')
    let res = await dispatch('fetchPage', payload)
    dispatch('fetchMenuItem')
    dispatch('fetchLogo')
    return res
  },
  initSettings ({ dispatch }, payload) {
  },
  setPreviewContent ({ commit }, payload) {
    commit(types.FETCH_PREVIEW_CONTENT, payload)
  },
  async fetchLogo ({ commit }) {
    const { data } = await axios.get(API_URL + '/api/logo')
    commit(types.FETCH_LOGO, data)
  },
  async fetchPublishedPages ({ commit }, payload) {
    const { data } = await axios.get(API_URL + '/api/published')
    commit(types.FETCH_PUBLISHED_PAGES, data)
  },
  async fetchArchivedPages ({ commit }, payload) {
    const { data } = await axios.get(API_URL + '/api/archived')
    commit(types.FETCH_ARCHIVED_PAGES, data)
  },
  async fetchAllContent ({ commit }, payload) {
    const { data } = await axios.get(API_URL + '/api/content')
    commit(types.FETCH_ALL_CONTENTS, data)
  },

  async fetchAllContactForm ({ commit }, payload) {
    const { data } = await axios.get(API_URL + '/api/contactform')
    commit(types.FETCH_ALL_CONTACTFORMS, data)
  },

  async fetchPage ({ commit, dispatch }, payload) {
    try {
      const { data } = await axios.get(API_URL + '/api/content/page', {params: {link: payload.page}})
      commit(types.FETCH_CURRENT_PAGES_CONTENTS, data)
      dispatch('fetchPageDetail', {page_id: data[0].page_id})
      dispatch('getPageVersion', {id: data[0].page_id})
      return true
    } catch (e) {
      commit(types.FETCH_CURRENT_PAGES_DETAILS, {})
      return false
    }
  },
  async fetchPageDetail ({ commit, dispatch }, payload) {
    const { data } = await axios.get(API_URL + '/api/page/' + payload.page_id)
    commit(types.FETCH_CURRENT_PAGES_DETAILS, data)
  },
  async fetchMenuItem ({ commit }) {
    const token = store.getters['auth/token']
    const { data } = await axios.get(API_URL + '/api/menu_item', {headers: {Authorization: `Bearer ${token}`}})
    commit(types.FETCH_ALL_MENUS, data)
  },
  async fetchAllTheme ({ commit }) {
    const token = store.getters['auth/token']
    const { data } = await axios.get(API_URL + '/api/extension', {headers: {Authorization: `Bearer ${token}`}, params: {type: 'theme'}})
    commit(types.FETCH_ALL_THEMES, data)
  },
  async fetchAllSystem ({ commit }) {
    const token = store.getters['auth/token']
    const { data } = await axios.get(API_URL + '/api/extension', {headers: {Authorization: `Bearer ${token}`}, params: {type: 'system'}})
    commit(types.FETCH_ALL_SYSTEM, data)
  },
  async fetchManifest ({ commit }, payload) {
    if (payload.theme) {
      const token = store.getters['auth/token']
      const { data } = await axios.get(API_URL + '/extension/theme/' + payload.theme, {headers: {Authorization: `Bearer ${token}`}})
      commit(types.FETCH_MANIFEST, data)
    } else {
      commit(types.FETCH_MANIFEST, {})
    }
  },
  async getPageVersion ({ commit }, payload) {
    const { data } = await axios.get(API_URL + '/api/page/' + payload.id + '/version')

    commit(types.FETCH_CURRENT_PAGES_VERSIONS, data)
  },
  async getPageDetail (state, payload) {
    const { data } = await axios.get(API_URL + '/api/page/' + payload)

    return data
  },
  async setPageVersion (state, payload) {
    const token = store.getters['auth/token']
    const { data } = await axios.put(API_URL + '/api/page/' + payload.id + '/revert', {}, {headers: {Authorization: `Bearer ${token}`}})

    return data
  },
  async deletePage ({ commit, dispatch }, payload) {
    await axios.delete(API_URL + '/api/page/' + payload)
    dispatch('init', '')
  }
}
