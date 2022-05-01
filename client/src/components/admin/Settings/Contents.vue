<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12">

      <div class="nav nav-tabs nav-justified" id="system-tab" role="tablist" aria-orientation="vertical">
        <a class="nav-item nav-link active" id="dq-accordion-contents-tab" data-toggle="pill" href="#dq-accordion-contents" role="tab" aria-controls="accordion-extensions" aria-selected="true">{{ $t('settings_contents') }}</a>
        <a class="nav-item nav-link" id="dq-accordion-contactforms-tab" data-toggle="pill" href="#dq-accordion-contactforms" role="tab" aria-controls="accordion-installed-extensions" aria-selected="true">{{ $t('settings_contents__contact_forms') }}</a>
      </div>
      <Card>
        <div class="tab-content">

          <!-- Tab "Contents" -->
          <div class="tab-pane fade show active p-3" role="tabpanel" id="dq-accordion-contents" aria-labelledby="dq-accordion-contents-tab">
            <DataQuery :data="contents" parent="dq-accordion-contents-tab" :query="contentQuery" v-if="$route.query.settings === 'contents'">
              <template slot="items" slot-scope="data" v-if="data.data">
                <div class="card">
                  <div class="card-header collapsed" data-toggle="collapse" :data-target="'#content-' + data.data.id" aria-expanded="true" :aria-controls="'content-' + data.data.id" @click="getContent(data.data.id)">
                    <h5 class="mb-0">
                      {{ data.data.name }}
                    </h5>
                  </div>
                  <div :id="'content-' + data.data.id" class="collapse" aria-labelledby="headingOne" data-parent="#dq-accordion-contents">
                    <div class="card-body">
                      <form @submit.prevent="saveContent" @keydown="form.onKeydown($event)">

                        <!-- Delete button -->
                        <div class="form-group row">
                          <span class="col-md-12 d-flex justify-content-start">
                            <v-button nativeType="button" :loading="processed" @click="removeContentInit(data.data)" class="btn btn-danger">
                              <span class="fa fa-trash" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__delete') }}
                            </v-button>
                          </span>
                        </div>
                        <hr>

                        <!-- Edit content -->
                        <div class="form-group row">
                          <span class="col-md-12">
                            <quill-editor :ref="'editor-' + data.data.id" v-model="form.content" v-if="form.id === data.data.id"></quill-editor>
                          </span>
                        </div>

                        <!-- Content name -->
                        <div class="form-group row">
                          <label class="col-md-4 col-form-label text-md-right">
                            {{ $t('settings_contents__name') }}
                          </label>
                          <span class="col-md-8">
                            <input v-model="form.name" type="text" name="name" class="form-control">
                          </span>
                        </div>

                        <!-- Save button -->
                        <div class="form-group row">
                          <span class="col-md-4 offset-md-4">
                            <v-button :loading="form.busy" class="btn btn-success">
                              {{ $t('button__save') }}
                            </v-button>
                          </span>
                        </div>

                        <!-- Used in pages -->
                        <div class="form-group row">
                          <label class="col-md-4 col-form-label text-md-right">{{ $t('settings_contents__used_in_page') }}</label>
                          <span class="col-md-8">
                            <h5>
                              <router-link v-for="(cl, key) in form.pages" :key="key" :to="{path: cl.link}" class="badge badge-dark mx-1">
                                {{ cl.name }}
                              </router-link>
                            </h5>
                          </span>
                        </div>

                      </form>
                    </div>
                  </div>
                </div>
              </template>
            </DataQuery>
          </div>

          <!-- Tab "Contact forms" -->
          <div class="tab-pane fade p-3" role="tabpanel" id="dq-accordion-contactforms" aria-labelledby="dq-accordion-contactforms-tab">

            <!-- Add contact form button -->
            <a class="dropdown-item" data-toggle="modal" data-target="#contactform-modal-index"><span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('settings_contents__new_contact_form') }}</a>
            <Newcontactform id="contactform-modal-index" />
            <br>

            <!-- Contact forms list -->
            <DataQuery :data="contactforms" parent="dq-accordion-contactforms" :query="contactFormQuery">
              <template slot="items" slot-scope="data" v-if="data.data">
                <div class="card">
                  <div class="card-header collapsed" data-toggle="collapse" :data-target="'#contactForm-' + data.keys" aria-expanded="true" :aria-controls="'contactForm-' + data.keys" >
                    <h5 class="mb-0">
                      {{ $t(data.data.name) }}
                    </h5>
                  </div>
                  <div :id="'contactForm-' + data.keys" class="collapse" aria-labelledby="headingOne" data-parent="#dq-accordion-contactforms">
                    <div class="card-body">

                      <div class="form-group row">
                        <span class="col-md-12">
                          <table class="table table-striped table-bordered" v-if="current.key === data.keys && current.data">
                            <tbody>
                              <tr v-for="(dt, key3, idx) in data.data" :key="idx" v-if="showVar.indexOf(key3) !== -1">
                                <td>
                                  {{ key3 }}
                                </td>
                                <td>
                                  {{ dt }}
                                </td>
                              </tr>
                            </tbody>
                          </table>
                        </span>
                      </div>

                      <template v-if="data.data.json_settings">

                        <!-- Delete button -->
                        <div class="form-group row">
                          <span class="col-md-12 d-flex justify-content-start">
                            <v-button :loading="processed" @click="deleteFormInit(data.data)" class="btn btn-danger">
                              <span class="fa fa-trash" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__delete') }}
                            </v-button>
                          </span>
                        </div>
                        <hr>

                        <!-- Recipient email address -->
                        <div class="form-group row" v-for="(st, key2) in data.data.json_settings" :key="key2">
                          <template v-if="st.type !== 'textarea'">
                            <label class="col-md-4 col-form-label text-md-right">
                              {{ $t('recipient_email_address') }}
                            </label>
                            <span class="col-md-8">
                              <input type="text" :name="st.name" :placeholder="st.placeholder" v-model="st.value" class="form-control">
                            </span>
                          </template>
                        </div>

                        <!-- Save button -->
                        <div class="form-group row">
                          <span class="col-md-4 offset-md-4">
                            <v-button :loading="processed" @click="changeSetting(data.data)" class="btn btn-success">
                              {{ $t('button__save') }}
                            </v-button>
                          </span>
                        </div>

                        <!-- Used in pages -->
                        <div class="form-group row">
                          <label class="col-md-4 col-form-label text-md-right">{{ $t('settings_contents__used_in_page') }}</label>
                          <span class="col-md-8">
                            <h5>
<!-- // Used for normal content, need to be adapted for contact-form
                              <router-link v-for="(cl, key) in form.pages" :key="key" :to="{path: cl.link}" class="badge badge-dark mx-1">
                                {{ cl.name }}
                              </router-link>
-->
                            </h5>
                          </span>
                        </div>

                      </template>
                    </div>
                  </div>
                </div>
              </template>
            </DataQuery>

          </div>

        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import _ from 'lodash'
import swal from 'sweetalert2'

import Newcontactform from '@/components/admin/Newcontactform.vue'
import axios from 'axios'
import Form from 'vform'

export default {
  name: 'admin-settings-contents',

  mounted () {
    this.$store.dispatch('page/fetchAllContactForm')
  },

  data: () => ({
    processed: false,
    template: {
      id: 0,
      name: '',
      content: '',
      pages: []
    },
    form: new Form({
      id: 0,
      name: '',
      content: '',
      pages: []
    }),
    contentQuery: {
      key: null,
      column: 'name,content',
      sort: null,
      sortColumn: null,
      limit: 10,
      offset: 0,
      current: 1
    },
    regexpPattern: /^.+@.+$/,
    editorOption: {
      theme: 'snow'
    },
    current: {
      key: null,
      data: {}
    },
    paginate: ['contactforms'],
    showVar: ['id'],
    contactFormQuery: {
      key: null,
      column: 'name,json_settings,id',
      sort: null,
      sortColumn: null,
      limit: 10,
      offset: 0,
      current: 1
    }
  }),

  computed: {
    ...mapGetters({
      contents: 'page/contents',
      token: 'auth/token',
      contactforms: 'page/contactforms',
      user: 'auth/user'
    })
  },

  components: {
    Newcontactform
  },

  methods: {
    getContent (v) {
      for (var i in this.contents) {
        if (this.contents[i].id === v) {
          var temp = this.template
          _.assign(this.form, this.template)
          _.assign(this.form, _.pick(this.contents[i], _.keys(temp)))
        }
      }
    },

    async saveContent () {
      if (!this.form.name || this.form.name.trim() === '') {
        this.$awn.alert(this.$t('settings_contents__empty_content_name'))
      } else {
        this.form.name = this.form.name.trim()
        try {
          await this.form.put(process.env.API_URL + '/api/content/' + this.form.id)
          this.$awn.success(this.form.name + ' - ' + this.$t('settings_contents__content_changes_saved'))
          await this.$store.dispatch('page/fetchAllContent')
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },

    // Ask confirmation before deleting Content
    async removeContentInit (v) {
      swal({
        title: this.$t('settings_contents__delete_content_warning_1'),
        text: this.$t('settings_contents__delete_content_warning_2') + v.name + this.$t('settings_contents__delete_content_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#d6c757d',
        confirmButtonText: this.$t('button__yes_delete'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.removeContent(v)
        }
      })
    },

    // Remove Content
    async removeContent (v) {
      this.processed = true
      try {
        await this.form.delete(process.env.API_URL + '/api/content/' + this.form.id, {
          headers: {Authorization: `Bearer ${this.token}`}
        })
        this.$awn.success(this.$t('settings_contents__content_deleted'))
        await this.$store.dispatch('page/fetchAllContent')
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },

    // Change contact form settings
    async changeSetting (v) {
      // Useless as long as one can't modify a form name
      // if (!v.name || v.name.trim() === '') {
      // this.$awn.alert(this.$t('settings_contents__empty_form_name'))
      if (!v.json_settings[0].value || !this.regexpPattern.test(v.json_settings[0].value.trim())) {
        this.$awn.alert(this.$t('recipient_email_address_is_invalid'))
      } else {
        this.processed = true
        v.name = v.name.trim()
        v.json_settings[0].value = v.json_settings[0].value.trim()
        try {
          let obj = JSON.parse(JSON.stringify(v))
          obj.json_settings = v.json_settings

          // Check if logged in user is creator
          // if (v.created_by === this.user.id) {
          if (v.created_by !== 'truc') {
            await axios.put(process.env.API_URL + '/api/contactform/' + v.id, obj, {
              headers: {Authorization: `Bearer ${this.token}`}
            })
            this.$awn.success(this.$t('settings_contents__contact_form_changes_saved'))
            await this.$store.dispatch('page/fetchAllContactForm')
            this.processed = false
          } else {
            this.$awn.alert(this.$t('settings_contents__cannot_edit_contact_form'))
          }
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
          this.processed = false
        }
      }
    },

    // Ask confirmation before deleting Contact form
    async deleteFormInit (v) {
      swal({
        title: this.$t('settings_contents__delete_contact_form_warning_1'),
        text: this.$t('settings_contents__delete_contact_form_warning_2') + v.name + this.$t('settings_contents__delete_contact_form_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#d6c757d',
        confirmButtonText: this.$t('button__yes_delete'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.deleteForm(v)
        }
      })
    },

    // Delete Contact form
    async deleteForm (v) {
      try {
        // if (v.created_by === this.user.id) {
        if (v.created_by !== 'truc') {
          await axios.delete(process.env.API_URL + '/api/contactform/' + v.id, {
            headers: {Authorization: `Bearer ${this.token}`}
          })
          this.$awn.success(this.$t('settings_contents__contact_form_deleted'))
          await this.$store.dispatch('page/fetchAllContactForm')
          this.processed = false
        } else {
          this.$awn.alert(this.$t('settings_contents__cannot_delete_contact_form'))
        }
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },

    beautifyDate (v) {
      let date = new Date(v * 1000)
      return ('0' + (date.getMonth() + 1)).slice(-2) + '/' + ('0' + date.getDate()).slice(-2) + '/' + date.getFullYear()
    },

    async saveThemeSettings () {
      let newThemeID = ''
      Object.keys(this.theme).forEach(key => {
        if (!isNaN(key) && this.theme[key]) {
          newThemeID = key
        }
      })
      try {
        const { data } = await axios.put(process.env.API_URL + '/api/theme?new=' + newThemeID, {}, {headers: {Authorization: `Bearer ${this.token}`}})
        this.$awn.success(this.$t(data.message))
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    }
  }
}
</script>

<style scoped>
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

.contactforms {
  margin-bottom: 50px;
}

</style>
