// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import Vue from 'vue'
import fontawesome from '@fortawesome/fontawesome'
import FontAwesomeIcon from '@fortawesome/vue-fontawesome'
import {
  faUser, faLock, faSignOutAlt, faCog
} from '@fortawesome/fontawesome-free-solid'

import { faGithub } from '@fortawesome/fontawesome-free-brands'

Vue.component('fa', FontAwesomeIcon)

// import { } from '@fortawesome/fontawesome-free-regular'

fontawesome.library.add(
  faUser, faLock, faSignOutAlt, faCog, faGithub
)
