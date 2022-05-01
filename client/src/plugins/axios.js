// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import axios from 'axios'
import store from '@/store'
import router from '@/router'
import swal from 'sweetalert2'
import i18n from '@/plugins/i18n.js'

// Request interceptor
axios.interceptors.request.use(request => {
  if (['post', 'put', 'delete'].indexOf(request.method) !== -1) {
    const token = store.getters['auth/token']
    if (token) {
      request.headers.common['Authorization'] = `Bearer ${token}`
    }

    // request.headers['X-Socket-Id'] = Echo.socketId()
  }

  const locale = store.getters['lang/locale']
  if (locale) {
    request.headers.common['Accept-Language'] = locale
  }

  return request
})

// Response interceptor
axios.interceptors.response.use(response => response, error => {
  const { status } = error.response

  if (status >= 500) {
    swal({
      type: 'error',
      title: i18n.t('error_occured_1'),
      text: i18n.t('error_occured_2'),
      reverseButtons: true,
      confirmButtonText: i18n.t('button__ok'),
      cancelButtonText: i18n.t('button__cancel')
    })
  }

  if (status === 401 && store.getters['auth/check']) {
    swal({
      type: 'warning',
      title: i18n.t('token_expired_1'),
      text: i18n.t('token_expired_2'),
      reverseButtons: true,
      confirmButtonText: i18n.t('button__ok'),
      cancelButtonText: i18n.t('button__cancel')
    }).then(async () => {
      await store.dispatch('auth/logout')

      router.push({ name: 'Login' })
    })
  }

  return Promise.reject(error)
})
