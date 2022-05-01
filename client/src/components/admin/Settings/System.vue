<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12">
      <div class="nav nav-tabs nav-justified" id="system-tab" role="tablist" aria-orientation="vertical">
        <a class="nav-item nav-link active" id="system-general-tab" data-toggle="pill" href="#system-general" role="tab" aria-controls="system-general" aria-selected="true">{{ $t('settings_system__main') }}</a>
        <a class="nav-item nav-link" id="system-mail-tab" data-toggle="pill" href="#system-mail" role="tab" aria-controls="system-mail" aria-selected="false">{{ $t('settings_system__smtp_server') }}</a>
        <a class="nav-item nav-link" id="system-file-tab" data-toggle="pill" href="#system-file" role="tab" aria-controls="system-file" aria-selected="false">{{ $t('settings_system__backups') }}</a>
      </div>
      <Card>
        <div class="tab-content" id="system-tabContent">

          <!-- Tab "General" -->
          <div class="tab-pane fade show active p-3" id="system-general" role="tabpanel" aria-labelledby="system-general-tab">

            <!-- Show/hide logo, menus and language selection in navigation bar -->
            <h5>
              {{ $t('settings_system__navigation_bar') }}
            </h5>
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input v-model="form.show_logo" type="checkbox" name="show_logo">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_system__show_logo') }}
              </label>
            </div>
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input v-model="form.show_menus" type="checkbox" name="show_menus">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_system__show_navigation_menus') }}
              </label>
            </div>
            <div class="form-group row">
              <span class="col-md-4 d-flex justify-content-end align-items-bottom">
                <input v-model="form.show_locales_menu" type="checkbox" name="show_locales_menu">
              </span>
              <label class="col-md-8 col-form-label">
                {{ $t('settings_system__show_language_menu') }}
              </label>
            </div>
            <div class="form-group row">
              <span class="col-md-4 offset-md-4">
                <v-button :loading="form.busy" @click="saveNavSettings" class="btn btn-success">
                  {{ $t('button__save') }}
                </v-button>
              </span>
            </div>
            <br>

            <!-- Application parameters -->
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__website_name') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.name" type="text" name="name" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__homepage') }}
              </label>
              <span class="col-md-8">
                <select v-model="form.home_page" name="home_page" class="form-control">
                  <option v-for="(pg, key) in allpages" :key="key" :value="pg.id">
                    {{ pg.name }}
                  </option>
                </select>
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__banned_attemps') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.banned_attemps" type="text" name="banned_attemps" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__lock_time') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.lock_time" type="text" name="lock_time" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__maximum_access_attemps') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.access_attemps" type="text" name="access_attemps" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__access_expired') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.access_expired" type="text" name="access_expired" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__secret_key') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.secret_key" type="text" name="secret_key" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__max_upload_size') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.max_size_upload" type="text" name="max_size_upload" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__favicon_url') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.favicon_url" type="text" name="favicon_url" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__logo_url') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.logo_url" type="text" name="logo_url" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <span class="col-md-4 offset-md-4">
                <v-button :loading="form.busy" @click="saveSetting" class="btn btn-success">
                  {{ $t('button__save') }}
                </v-button>
              </span>
            </div>

          </div>

          <!-- Tab "SMTP server" -->
          <div class="tab-pane fade p-3" id="system-mail" role="tabpanel" aria-labelledby="system-mail-tab">

            <!-- SMTP parameters -->
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__smtp_host') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.smtp_host" type="text" name="smtp_host" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__smtp_port') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.smtp_port" type="text" name="smtp_port" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__smtp_user') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.smtp_user" type="text" name="smtp_user" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__smtp_password') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.smtp_password" type="password" name="smtp_password" class="form-control">
              </span>
            </div>
            <br>
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('settings_system__sender_email_address') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.email_sender" type="text" name="email_sender" class="form-control">
              </span>
            </div>
            <div class="form-group row">
              <span class="col-md-4 offset-md-4">
                <v-button :loading="form.busy" @click="saveSetting" class="btn btn-success">
                  {{ $t('button__save') }}
                </v-button>
              </span>
            </div>

          </div>

          <!-- Tab "Backups" -->
          <div class="tab-pane fade p-3" id="system-file" role="tabpanel" aria-labelledby="system-file-tab">
            <div class="form-group row">
              <span class="col-md-12 d-flex justify-content-end">
                <v-button :loading="processed" @click="setBackup" class="btn btn-primary">
                  {{ $t('button__backup_system_and_data') }}
                </v-button>
              </span>
            </div>
            <br>
            <div class="form-group row">
              <span class="col-md-12">
                <table class="table table-striped table-bordered">
                  <thead>
                    <tr>
                      <th>{{ $t('settings_system__backup_file_name') }}</th>
                      <th>{{ $t('settings_system__created_on') }}</th>
                      <th></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(st, key, idx) in backups" :key="idx">
                      <td>{{ st.name }}</td>
                      <td>{{ st.created_on }}</td>
                      <td>
                        <v-button :loading="processed" @click="restoreSystemInit(st)" class="btn btn-danger">
                          <span class="fa fa-undo" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__system_restore') }}
                        </v-button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </span>
            </div>
          </div>

        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import axios from 'axios'
import swal from 'sweetalert2'
import _ from 'lodash'
export default {
  name: 'admin-settings-system',

  mounted () {
    this.checkUser()
    this.fetchNavSettings()
    this.$nextTick(() => {
      if (this.$route.query.settings === 'system') {
        this.fetchSetting()
        this.getBackup()
        this.getAllPages()
      }
    })
  },

  data () {
    return {
      regexpPattern: /^.+@.+$/,
      processed: false,
      form: new Form({
        show_logo: false,
        show_menus: false,
        show_locales_menu: false,
        smtp_host: '',
        smtp_port: '',
        smtp_user: '',
        smtp_password: '',
        email_sender: '',
        name: '',
        home_page: '',
        banned_attemps: '',
        lock_time: '',
        access_attemps: '',
        access_expired: '',
        secret_key: '',
        max_size_upload: '',
        favicon_url: '',
        logo_url: ''
      }),
      backups: [],
      allpages: []
    }
  },

  computed: mapGetters({
    token: 'auth/token',
    user: 'auth/user'
  }),

  methods: {

    // Navigation settings methods
    async fetchNavSettings () {
      try {
        const data = await this.$store.dispatch('auth/fetchSetting')
        _.assign(this.form, data)
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },
    async saveNavSettings () {
      try {
        await this.form.put(process.env.API_URL + '/api/setting', {}, {headers: {Authorization: `Bearer ${this.token}`}})
        this.$awn.success(this.$t('settings_system__navigation_bar_settings_saved'))
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
      if (this.form.show_logo) {
        document.querySelector('header a.navbar-brand').style.display = 'flex'
      } else {
        document.querySelector('header a.navbar-brand').style.display = 'none'
      }
      if (this.form.show_locales_menu) {
        document.querySelector('#navbarSupportedContent>ul.navbar-nav.ml-auto:nth-child(2)').style.display = 'flex'
      } else {
        console.log(document.querySelector('#navbarSupportedContent>ul.navbar-nav.ml-auto:nth-child(2)'))
        document.querySelector('#navbarSupportedContent>ul.navbar-nav.ml-auto:nth-child(2)').style.display = 'none'
      }
      if (this.form.show_logo || this.form.show_menus || this.form.show_locales_menu) {
        document.querySelector('header').style.display = 'flex'
      } else {
        document.querySelector('header').style.display = 'none'
      }
      this.fetchNavSettings()
      if (this.form.show_menus) {
        document.querySelector('#navbarSupportedContent>div:first-child').style.display = 'flex'
        document.querySelector('ul.sidebar-menu i.fa.fa-bars').parentNode.parentNode.style.display = 'flex'
      } else {
        document.querySelector('#navbarSupportedContent>div:first-child').style.display = 'none'
        document.querySelector('ul.sidebar-menu i.fa.fa-bars').parentNode.parentNode.style.display = 'none'
      }
    },

    // Set form into default value
    // initForm () {
    //  this.form = new Form({
    //    access_level: 1,
    //    name: '',
    //    title: '',
    //    link: '',
    //    keywords: '',
    //    description: ''
    //  })
    // },

    // Fetch existing settings
    async fetchSetting () {
      try {
        const data = await this.$store.dispatch('auth/fetchSetting')

        _.assign(this.form, data)
      } catch (e) {}
    },

    // Get pages list
    async getAllPages () {
      this.processed = true
      try {
        const { data } = await axios.get(process.env.API_URL + '/api/listallpages', {headers: {Authorization: `Bearer ${this.token}`}})

        this.$set(this, 'allpages', data)
        this.processed = false
      } catch (e) {
        this.processed = false
      }
    },

    // Save settings
    async saveSetting () {
      if (!this.form.smtp_host || this.form.smtp_host.trim() === '') {
        this.$awn.alert(this.$t('settings_system__empty_SMPT_server_address'))
      } else if (!this.form.smtp_user || !this.regexpPattern.test(this.form.smtp_user.trim())) {
        this.$awn.alert(this.$t('settings_accounts__invalid_email_address'))
      } else if (!this.form.smtp_password || this.form.smtp_password.trim() === '') {
        this.$awn.alert(this.$t('settings_accounts__invalid_password'))
      } else if (!this.form.email_sender || !this.regexpPattern.test(this.form.email_sender.trim())) {
        this.$awn.alert(this.$t('settings_system__invalid_sender_email_address'))
      } else {
        try {
          this.form.smtp_host = this.form.smtp_host.trim()
          this.form.smtp_port = parseInt(this.form.smtp_port)
          this.form.smtp_user = this.form.smtp_user.trim()
          this.form.email_sender = this.form.email_sender.trim()
          this.form.banned_attemps = parseInt(this.form.banned_attemps)
          this.form.lock_time = parseInt(this.form.lock_time)
          this.form.access_attemps = parseInt(this.form.access_attemps)
          this.form.access_expired = parseInt(this.form.access_expired)
          this.form.max_size_upload = parseInt(this.form.max_size_upload)
          await this.form.put(process.env.API_URL + '/api/setting', {}, {headers: {Authorization: `Bearer ${this.token}`}})
          this.fetchSetting()
          this.$awn.success(this.$t('settings_system__settings_saved'))
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },

    // Check user access_level and push 404 if user trying to access this page.
    // TODO: should be done by Go code, not Vue???
    async checkUser () {
      if (this.user.access_level !== 10) {
        this.$router.push('/404')
      }
    },

    // Get backup list
    async getBackup () {
      this.processed = true
      try {
        const { data } = await axios.get(process.env.API_URL + '/api/file/backup_list', {headers: {Authorization: `Bearer ${this.token}`}})

        this.$set(this, 'backups', data)
        this.processed = false
      } catch (e) {
        this.processed = false
      }
    },

    // Make a backup
    async setBackup () {
      this.processed = true
      try {
        await axios.get(process.env.API_URL + '/api/file/backup', {headers: {Authorization: `Bearer ${this.token}`}})

        this.$awn.success(this.$t('settings_system__backup_done'))
        this.getBackup()
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },

    // Restore system initial
    async restoreSystemInit (v) {
      swal({
        title: this.$t('button__system_restore'),
        text: this.$t('settings_system__restore_warning_1') + v.created_on + this.$t('settings_system__restore_warning_2'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_restore'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.restoreSystem(v)
        }
      })
    },

    // Restore system
    async restoreSystem (v) {
      this.processed = true
      try {
        const { data } = await axios.get(process.env.API_URL + '/api/system/revert', {
          headers: {Authorization: `Bearer ${this.token}`},
          params: {
            filename: v.name
          }
        })

        this.$awn.success(this.$t(data.message))
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    }
  }
}
</script>

<style scoped>
.card-body {
  padding-top: 0
}
</style>
