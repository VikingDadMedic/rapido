<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>,Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12" id="accordion">
      <div class="nav nav-tabs nav-justified" id="system-tab" role="tablist" aria-orientation="vertical">
        <a class="nav-item nav-link active" id="accordion-users-tab" data-toggle="pill" href="#accordion-users" role="tab" aria-controls="accordion-users" aria-selected="true">{{ $t('settings_accounts__accounts') }}</a>
        <a v-if="user.access_level === 10" class="nav-item nav-link" id="accordion-add-user-tab" data-toggle="pill" href="#accordion-add-user" role="tab" aria-controls="accordion-add-user" aria-selected="true">{{ $t('settings_accounts__create_account') }}</a>
      </div>
      <Card>
        <div class="tab-content">

          <!-- Tab "Accounts" -->
          <div class="tab-pane fade show active p-3" id="accordion-users" role="tabpanel" aria-labelledby="accordion-users-tab">
            <div class="card" v-for="(us, key) in users" :key="key">

              <!-- Account mail address + load data -->
              <div class="card-header collapsed" data-toggle="collapse" :data-target="'#user-' + us.id" aria-expanded="true" :aria-controls="'user-' + us.id" v-on:click="getUser(us.id)">
                <h5 class="mb-0">
                    {{ us.email_address }}&nbsp;&nbsp;<span v-if="us.active !== 1">{{ $t('settings_accounts__unactivated')  }}</span>
                </h5>
              </div>

              <!-- Forms -->
<!--              <div v-if="us.id === user.id" :id="'user-' + us.id" class ="collapse show" aria-labelledby="headingOne" data-parent="#accordion-users" v-on:click="getUser(us.id)"> Tried to show by default logged in user' account but click event no more happens -->
              <div :id="'user-' + us.id" class ="collapse" aria-labelledby="headingOne" data-parent="#accordion-users">
                <div class="card-body">
                  <form @submit.prevent="editUser" @keydown="form.onKeydown($event)">

                    <!-- Delete button -->
                    <div class="form-group row">
                      <span class="col-md-12 d-flex justify-content-start">
                        <v-button nativeType="button" :loading="processed" @click="removeUserInit(us)" class="flex-grow-0 btn btn-danger">
                          <span class="fa fa-trash" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__delete') }}
                        </v-button>
                      </span>
                    </div>
                    <hr>

                    <!-- Access Level -->
                    <div class="form-group row">
                      <label class="col-md-4 col-form-label text-md-right">
                        {{ $t('settings_accounts__access_level') }}
                      </label>
                      <span class="col-md-3">
                        <select v-model="form.access_level" id="access_level" name="access_level" class="form-control">
                          <option v-for="(al, key) in user_level" :key="key" :value="al.value">
                            {{ $t(al.name) }}
                          </option>
                        </select>
                      </span>
                    </div>

                    <!-- Email Address -->
                    <div class="form-group row">
                      <label class="col-md-4 col-form-label text-md-right">
                        {{ $t('settings_accounts__email_address') }}
                      </label>
                      <span class="col-md-8">
                        <input v-model="form.email_address" type="text" name="email_address" class="form-control">
                      </span>
                    </div>

                    <!-- Submit Button -->
                    <div class="form-group row">
                      <span class="col-md-8 offset-md-4">
                        <v-button :loading="form.busy" class="flex-grow-0 btn btn-success">
                          {{ $t('button__save') }}
                        </v-button>
                      </span>
                    </div>
                  </form>
                  <form @submit.prevent="changePassword" @keydown="form.onKeydown($event)">

                    <!-- Password -->
                    <div class="form-group row">
                      <label class="col-md-4 col-form-label text-md-right">
                        {{ $t('settings_accounts__password') }}
                      </label>
                      <span class="col-md-8">
                        <input v-model="form2.password" type="password" name="password" class="form-control">
                      </span>
                    </div>

                    <!-- Confirmation password -->
                    <div class="form-group row">
                      <label class="col-md-4 col-form-label text-md-right">
                        {{ $t('settings_accounts__confirm_password') }}
                      </label>
                      <span class="col-md-8">
                        <input v-model="form2.confirmation_password" type="password" name="confirmation_password" class="form-control">
                      </span>
                    </div>

                    <!-- Submit Button -->
                    <div class="form-group row">
                      <span class="col-md-4 offset-md-4 d-flex justify-content-start">
                        <v-button :loading="form2.busy" class="flex-grow-0 btn btn-success">
                          {{ $t('button__change_password') }}
                        </v-button>
                      </span>
                    </div>

                  </form>
                </div>
              </div>

            </div>
          </div>

          <!-- Tab "Create account" -->
          <div class="tab-pane fade p-3" id="accordion-add-user" role="tabpanel" aria-labelledby="accordion-add-user-tab">
            <form @submit.prevent="addUser" @keydown="form3.onKeydown($event)">

              <!-- Access Level -->
              <div class="form-group row">
                <label class="col-md-4 col-form-label text-md-right">
                  {{ $t('settings_accounts__access_level') }}
                </label>
                <span class="col-md-3">
                  <select v-model="form3.access_level" id="access_level" name="access_level" class="form-control">
                    <option v-for="(al, key) in user_level" :key="key" :value="al.value">
                      {{ $t(al.name) }}
                    </option>
                  </select>
                </span>
              </div>

              <!-- Email -->
              <div class="form-group row">
                <label class="col-md-4 col-form-label text-md-right">
                  {{ $t('settings_accounts__email_address') }}
                </label>
                <span class="col-md-8">
                  <input v-model="form3.email_address" type="text" name="email_address" class="form-control">
                </span>
              </div>

              <!-- Password -->
              <div class="form-group row">
                <label class="col-md-4 col-form-label text-md-right">
                  {{ $t('settings_accounts__password') }}
                </label>
                <span class="col-md-8">
                  <input v-model="form3.password" type="password" name="password" class="form-control">
                </span>
              </div>

              <!-- Confirmation password -->
              <div class="form-group row">
                <label class="col-md-4 col-form-label text-md-right">
                  {{ $t('settings_accounts__confirm_password') }}
                </label>
                <span class="col-md-8">
                  <input v-model="form3.confirmation_password" type="password" name="confirmation_password" class="form-control">
                </span>
              </div>

              <!-- Submit Button -->
              <div class="form-group row">
                <span class="col-md-8 offset-md-4">
                  <v-button :loading="form3.busy" class="flex-grow-0 btn btn-success">
                    {{ $t('button__save') }}
                  </v-button>
                </span>
              </div>

            </form>
          </div>

        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import _ from 'lodash'
import axios from 'axios'
import swal from 'sweetalert2'
import Form from 'vform'
export default {
  name: 'admin-settings-users',

  mounted () {
    this.$store.dispatch('auth/fetchAllUser')
  },

  data: () => ({
    processed: false,
    regexpPattern: /^.+@.+$/,
    form: new Form({
      id: 0,
      email_address: '',
      password: '',
      access_level: ''
    }),
    form2: new Form({
      id: 0,
      password: '',
      confirmation_password: ''
    }),
    form3: new Form({
      email_address: '',
      password: '',
      access_level: ''
    })
  }),

  computed: mapGetters({
    token: 'auth/token',
    user: 'auth/user',
    user_level: 'auth/user_level',
    users: 'auth/users'
  }),

  methods: {
    getUser (v) {
      for (var i in this.users) {
        if (this.users[i].id === v) {
          var temp = this.form
          this.form = new Form({})
          _.assign(this.form, _.pick(this.users[i], _.keys(temp)))
          var temp2 = this.form2
          this.form2 = new Form({})
          _.assign(this.form2, _.pick(this.users[i], _.keys(temp2)))
        }
      }
    },
    async editUser () {
      if (!this.form.email_address || !this.regexpPattern.test(this.form.email_address.trim())) {
        this.$awn.alert(this.$t('settings_accounts__invalid_email_address'))
      } else {
        this.form.email_address = this.form.email_address.trim()
        try {
          await this.form.put(process.env.API_URL + '/api/account/' + this.form.id)
          this.$awn.success(this.$t('settings_accounts__account_changes_saved'))
          await this.$store.dispatch('auth/fetchAllUser')
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },
    async changePassword () {
      if (!this.form2.password || this.form2.password.trim() === '' || !this.form2.confirmation_password || this.form2.confirmation_password.trim() === '') {
        this.$awn.alert(this.$t('settings_accounts__invalid_password'))
      } else if (this.form2.password !== this.form2.confirmation_password) {
        this.$awn.alert(this.$t('settings_accounts__passwords_do_not_match'))
      } else {
        try {
          const { data } = await this.form2.put(process.env.API_URL + '/api/change_password/' + this.form2.id)
          this.$awn.success(this.$t(data.message))
          await this.$store.dispatch('auth/fetchAllUser')
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },
    // Delete user account initial
    async removeUserInit (v) {
      swal({
        title: this.$t('settings_accounts__delete_account_warning_1'),
        text: this.$t('settings_accounts__delete_account_warning_2') + v.email_address + this.$t('settings_accounts__delete_account_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_delete'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.removeUser(v)
        }
      })
    },
    // Delete user account
    async removeUser (v) {
      this.processed = true
      try {
        await axios.delete(process.env.API_URL + '/api/account/' + v.id, {
          headers: {Authorization: `Bearer ${this.token}`}
        })

        this.$awn.success(this.$t('settings_accounts__account_deleted'))
        await this.$store.dispatch('auth/fetchAllUser')
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    async addUser () {
      this.processed = true
      if (!this.form3.access_level) {
        this.$awn.alert(this.$t('settings_accounts__invalid_access_level'))
      } else if (!this.form3.email_address || !this.regexpPattern.test(this.form3.email_address.trim())) {
        this.$awn.alert(this.$t('settings_accounts__invalid_email_address'))
      } else if (!this.form3.password || this.form3.password.trim() === '' || !this.form3.confirmation_password || this.form3.confirmation_password.trim() === '') {
        this.$awn.alert(this.$t('settings_accounts__invalid_password'))
      } else if (this.form3.password !== this.form3.confirmation_password) {
        this.$awn.alert(this.$t('settings_accounts__passwords_do_not_match'))
      } else {
        this.form3.email_address = this.form3.email_address.trim()
        try {
          await this.form3.post(process.env.API_URL + '/api/account')
          this.$awn.success(this.$t('settings_accounts__account_created'))
          await this.$store.dispatch('auth/fetchAllUser')
          this.processed = false
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
          this.processed = false
        }
      }
    }
  }
}
</script>

<style scoped>
#access_level{
  width: 160%;
}

.card-header > h5 {
  cursor: pointer;
}
.card-header > h5:after {
  font-family: 'ForkAwesome';
  content: "\f106";
  float: right;
}
.card-header.collapsed > h5:after {
  content: "\f107";
}
</style>
