// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import Index from '@/pages/Index.vue'
import Login from '@/pages/auth/Login.vue'
import Register from '@/pages/auth/Register.vue'
import Email from '@/pages/auth/password/Email.vue'
import Reset from '@/pages/auth/password/Reset.vue'
import Error404 from '@/pages/errors/404.vue'

export default [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  // Guest routes.
  ...middleware('guest', [
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/forgot',
      name: 'Email',
      component: Email
    },
    {
      path: '/forgot_password',
      name: 'Reset',
      component: Reset
    },
    {
      path: '/user_activation',
      name: 'Activation',
      // component: Activation
      component: Login
    }
  ]),
  {
    path: '/404',
    name: 'not-found',
    component: Error404
  },
  {
    path: '*',
    name: 'Page',
    component: Index
  }
]

/**
 * @param  {String|Function} middleware
 * @param  {Array} routes
 * @return {Array}
 */
function middleware (middleware, routes) {
  routes.forEach(route =>
    (route.middleware || (route.middleware = [])).unshift(middleware)
  )

  return routes
}
