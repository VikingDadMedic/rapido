// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


import Vue from 'vue'
import { HasError, AlertError, AlertSuccess } from 'vform'
import VueAWN from 'vue-awesome-notifications'
import VueQuillEditor from 'vue-quill-editor'
import Switches from 'vue-switches'
import VueClipboard from 'vue-clipboard2'
import VeeValidate from 'vee-validate'
import Quill from 'quill'
import { ImageUpload } from '@/plugins/image-upload.js'
import ImageResize from 'quill-image-resize-module'
import VideoResize from 'quill-video-resize-module'

import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

require('./validation.js')

Quill.register('modules/imageUpload', ImageUpload)
Quill.register('modules/imageResize', ImageResize)
Quill.register('modules/VideoResize', VideoResize)

let optionsAwn = {
  position: 'top-right'
}
let optionsQuill = {
  theme: 'snow',
  modules: {
    imageResize: {
      modules: [ 'Resize', 'DisplaySize', 'Toolbar' ]
    },
    VideoResize: {},
    toolbar: {
      container: [
        ['bold', 'italic', 'underline', 'strike'],
        ['blockquote', 'code-block'],
        [{ 'header': '1' }, { 'header': '2' }],
        [{ 'list': 'ordered' }, { 'list': 'bullet' }],
        [{ 'script': 'sub' }, { 'script': 'super' }],
        [{ 'indent': '-1' }, { 'indent': '+1' }],
        [{ 'direction': 'rtl' }],
        [{ 'size': [] }],
        [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
        [{ 'font': [] }],
        [{ 'color': [] }, { 'background': [] }],
        [{ 'align': [] }],
        ['clean'],
        ['link', {'customimage': [0, 1, 2]}, 'video']
      ],
      handlers: {
        'customimage' (value) {
          if (value === '1') {
            ImageUpload.selectLocalImage()
          }
        }
      },
      ImageResize: {
        modules: [ 'Resize', 'DisplaySize', 'Toolbar' ]
      }
    },
    imageUpload: {
      url: process.env.API_URL + '/api/file/media', // server url
      method: 'POST', // change query method, default 'POST'
      headers: {'Content-Type': 'application/x-www-form-urlencoded'}, // add custom headers, example { token: 'your-token'}
      // personalize successful callback and call next function to insert new url to the editor
      callbackOK: (serverResponse, next) => {
        next(serverResponse)
      },
      // personalize failed callback
      // callbackKO: (serverError) => {
      //   console.log(serverError)
      // },
      checkBeforeSend: (file, next) => {
        next(file)
      }
    }
  }
}

Vue.use(VueAWN, optionsAwn)
Vue.use(VueQuillEditor, optionsQuill)
Vue.use(VueClipboard)

Vue.config.productionTip = process.env.NODE_ENV === 'production'

Vue.use(VeeValidate)

Vue.component(HasError.name, HasError)
Vue.component(AlertError.name, AlertError)
Vue.component(AlertSuccess.name, AlertSuccess)
Vue.component(Switches.name, Switches)

// Load global components dynamically
const requireContext = require.context('./global', false, /.*\.(js|vue)$/)
requireContext.keys().forEach(file => {
  const Component = requireContext(file).default

  if (Component.name) {
    Vue.component(Component.name, Component)
  }
})

Vue.prototype.$toCamelCase = (str) => {
  return str.replace(/\W+(.)/g, function (match, chr) {
    return chr.toUpperCase()
  })
}

Vue.prototype.$toTitleCase = (str) => {
  return str.replace(/\w\S*/g, (txt) => {
    return txt.charAt(0).toUpperCase() + txt.substr(1)
  })
}
