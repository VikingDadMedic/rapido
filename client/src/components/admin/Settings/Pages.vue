<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>,Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12">
      <div class="nav nav-tabs nav-justified" id="system-tab" role="tablist" aria-orientation="vertical">
        <a class="nav-item nav-link active" id="accordion-published-tab" data-toggle="pill" href="#accordion-published" role="tab" aria-controls="accordion-published">{{ $t('settings_pages__visible_pages') }}</a>
        <a class="nav-item nav-link" id="accordion-archived-tab" data-toggle="pill" href="#accordion-archived" role="tab" aria-controls="accordion-archived">{{ $t('settings_pages__archived_pages') }}</a>
      </div>
      <Card>
        <div class="tab-content">
          <div class="tab-pane fade show active p-3" role="tabpanel" id="accordion-published" aria-labelledby="accordion-published-tab">
            <DataQuery :data="published_pages" parent="dq-accordion-published" :query="publishedQuery" v-if="$route.query.settings === 'pages'">
              <template slot="items" slot-scope="data" v-if="data.data">
                <div class="card">
                  <div class="card-header collapsed" data-toggle="collapse" :data-target="'#published-' + data.data.id" aria-expanded="true" :aria-controls="'published-' + data.data.id" @click="getPublishedPage(data.data.id)">
                    <h5 class="mb-0">
                      {{ data.data.name }} {{ data.data.link ? '(' + data.data.link + ')' : '' }}
                    </h5>
                  </div>
                  <div :id="'published-' + data.data.id" class="collapse" data-parent="#dq-accordion-published">
                    <div class="card-body">
                      <div class="form-group row">

                        <!-- Archive page button -->
                        <span class="col-md-4 offset-md-4 d-flex justify-content-center">
                          <v-button nativeType="button" :loading="processed" @click="archivePageInit(data.data)" class="btn btn-dark">
                            <span class="fa fa-moon-o" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__archive') }}
                          </v-button>
                        </span>

                        <!-- See page button -->
                        <span class="col-md-4 d-flex justify-content-end">
                          <router-link :to="{path: published_pages[data.keys].link ? (published_pages[data.keys].link.substring(0) === '/' ? published_pages[data.keys].link : '/' + published_pages[data.keys].link) : '/'}" tag="button" class="btn btn-primary">
                            <span class="fa fa-eye" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__see_page') }}
                          </router-link>
                        </span>

                      </div>
                      <hr>
                      <div class="form-group row">

                        <!-- Change published page settings -->
                        <span class="col-md-6">
                          <form @submit.prevent="savePublishedPage" @keydown="formpublished1.onKeydown($event)" data-vv-scope="publishedpagesettings">
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__name') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formpublished1.name" type="text" name="name">
                              </span>
                            </div>
                            <hr>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__title') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formpublished1.title" type="text" name="title" class="form-control">
                              </span>
                            </div>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__keywords') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formpublished1.keywords" type="text" name="keywords" class="form-control">
                              </span>
                            </div>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__description') }}
                              </label>
                              <span class="col-md-8">
                                <textarea v-model="formpublished1.description" name="description" class="form-control"></textarea>
                              </span>
                            </div>
                            <div class="form-group row">
                              <span class="col-md-4 offset-md-4 d-flex justify-content-begin">
                                <v-button :loading="formpublished1.busy" class="btn btn-success">
                                  {{ $t('button__save') }}
                                </v-button>
                              </span>
                            </div>
                          </form>
                        </span>

                        <!-- Displays page preview -->
                        <span class="col-md-6">
                          <div class="iframe-box" :ref="'ifbox-' + data.data.id">
                            <iframe :src="'/' + (data.data.link || '') + '?preview=true'" width="1366" height="700" :style="'transform: scale(' + $refs['ifbox-' + data.data.id].clientWidth / 1366 + ')'" v-if="formpublished1.id === data.data.id && $refs['ifbox-' + data.data.id]"></iframe>
                          </div>
                        </span>

                      </div>
                      <div class="form-group row">

                        <!-- Add author to published page -->
                        <span class="col-md-6">
                          <form @submit.prevent="addAuthorPublishedPage" @keydown="formpublished2.onKeydown($event)" data-vv-scope="publishedpageauthor">
                            <h5>
                              {{ $t('settings_pages__add_author') }}
                            </h5>
                            <div class="input-group">
                              <span class="input-group-prepend">
                                <span class="input-group-text">
                                  {{ $t('settings_accounts__email_address') }}
                                </span>
                              </span>
                              <input v-model="formpublished2.email" type="text" name="email" class="form-control">
                              <v-button :loading="formpublished1.busy" class="btn btn-success">
                                {{ $t('button__add') }}
                              </v-button>
                            </div>
                          </form>
                        </span>

                        <!-- List author(s) -->
                        <span class="col-md-6">
                          <h5>
                            {{ $t('settings_pages__authors') }}
                          </h5>
                          <h5 v-if="details.authors">
                            <span class="badge badge-danger font-weight-normal mx-1 mb-1 p-2" v-for="(cl, index) in details.authors" :key="index">
                              {{ cl }}&nbsp;<a class="ml-2" @click="removeAuthorInitPublishedPage(cl)"><span class="fa fa-times-circle"></span></a>
                            </span>
                          </h5>
                          <h5 class="text-center" v-else>
                            {{ $t('settings_pages__no_author') }}
                          </h5>
                        </span>

                      </div>
                    </div>
                  </div>
                </div>
              </template>
            </DataQuery>
          </div>
          <div class="tab-pane fade p-3" role="tabpanel" id="accordion-archived" aria-labelledby="accordion-archived-tab">
            <DataQuery :data="archived_pages" parent="dq-accordion-archived" :query="archivedQuery" v-if="$route.query.settings === 'pages'">
              <template slot="items" slot-scope="data" v-if="data.data">
                <div class="card">
                  <div class="card-header collapsed" data-toggle="collapse" :data-target="'#archived-' + data.data.id" aria-expanded="true" :aria-controls="'archived-' + data.data.id" @click="getArchivedPage(data.data.id)">
                    <h5 class="mb-0">
                      {{ data.data.name }} {{ data.data.link ? '(' + data.data.link + ')' : '' }}
                    </h5>
                  </div>
                  <div :id="'archived-' + data.data.id" class="collapse" data-parent="#dq-accordion-archived">
                    <div class="card-body">
                      <div class="form-group row">

                        <!-- Delete button -->
                        <span class="col-md-4 d-flex justify-content-start">
                          <v-button :loading="processed" @click="deletePageInit(data.data)" class="btn btn-danger">
                            <span class="fa fa-trash" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__delete') }}
                          </v-button>
                        </span>

                        <!-- Publish button -->
                        <span class="col-md-4 d-flex justify-content-center">
                          <v-button :loading="processed" @click="publishPageInit(data.data)" class="btn btn-warning">
                            <span class="fa fa-sun-o" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__publish') }}
                          </v-button>
                        </span>

                        <!-- See page button -->
                        <span class="col-md-4 d-flex justify-content-end">
                          <router-link :to="{path: archived_pages[data.keys].link ? (archived_pages[data.keys].link.substring(0) === '/' ? archived_pages[data.keys].link : '/' + archived_pages[data.keys].link) : '/'}" tag="button" class="btn btn-primary">
                            <span class="fa fa-eye" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__see_page') }}
                          </router-link>
                        </span>

                      </div>
                      <hr>
                      <div class="form-group row">

                        <!-- Change archived page settings -->
                        <span class="col-md-6">
                          <form @submit.prevent="saveArchivedPage" @keydown="formarchived1.onKeydown($event)" data-vv-scope="archivedpagesettings">
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__name') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formarchived1.name" type="text" name="name">
                              </span>
                            </div>
                            <hr>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__title') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formarchived1.title" type="text" name="title" class="form-control">
                              </span>
                            </div>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__keywords') }}
                              </label>
                              <span class="col-md-8">
                                <input v-model="formarchived1.keywords" type="text" name="keywords" class="form-control">
                              </span>
                            </div>
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right">
                                {{ $t('settings_pages__description') }}
                              </label>
                              <span class="col-md-8">
                                <textarea v-model="formarchived1.description" name="description" class="form-control"></textarea>
                              </span>
                            </div>
                            <div class="form-group row">
                              <span class="col-md-4 offset-md-4 d-flex justify-content-begin">
                                <v-button :loading="formarchived1.busy" class="btn btn-success">
                                  {{ $t('button__save') }}
                                </v-button>
                              </span>
                            </div>
                          </form>
                        </span>

                        <!-- Displays page preview -->
                        <span class="col-md-6">
                          <div class="iframe-box" :ref="'ifbox-' + data.data.id">
                            <iframe :src="'/' + (data.data.link || '') + '?preview=true'" width="1366" height="700" :style="'transform: scale(' + $refs['ifbox-' + data.data.id].clientWidth / 1366 + ')'" v-if="formarchived1.id === data.data.id && $refs['ifbox-' + data.data.id]"></iframe>
                          </div>
                        </span>

                      </div>
                      <div class="form-group row">

                        <!-- Add author to archived page -->
                        <span class="col-md-6">
                          <form @submit.prevent="addAuthorArchivedPage" @keydown="formarchived2.onKeydown($event)" data-vv-scope="archivedpageauthor">
                            <h5>
                              {{ $t('settings_pages__add_author') }}
                            </h5>
                            <div class="input-group">
                              <span class="input-group-prepend">
                                <span class="input-group-text">
                                  {{ $t('settings_accounts__email_address') }}
                                </span>
                              </span>
                              <input v-model="formarchived2.email" type="text" name="email" class="form-control">
                              <v-button :loading="formarchived1.busy" class="btn btn-success">
                                {{ $t('button__add') }}
                              </v-button>
                            </div>
                          </form>
                        </span>

                        <!-- List author(s) -->
                        <span class="col-md-6">
                          <h5>
                            {{ $t('settings_pages__authors') }}
                          </h5>
                          <h5 v-if="details.authors">
                            <span class="badge badge-danger font-weight-normal mx-1 mb-1 p-2" v-for="(cl, index) in details.authors" :key="index">
                              {{ cl }}&nbsp;<a class="ml-2" @click="removeAuthorInitArchivedPage(cl)"><span class="fa fa-times-circle"></span></a>
                            </span>
                          </h5>
                          <h5 class="text-center" v-else>
                            {{ $t('settings_pages__no_author') }}
                          </h5>
                        </span>

                      </div>
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
import swal from 'sweetalert2'
import _ from 'lodash'
import axios from 'axios'
import Form from 'vform'
export default {
  name: 'admin-settings-pages',

  async mounted () {
  },

  data: () => ({
    processed: false,
    regexpPattern: /^.+@.+$/,
    template: {
      id: 0,
      published: 0,
      authors: [],
      name: '',
      title: '',
      link: '',
      keywords: '',
      description: ''
    },
    formpublished1: new Form({
      id: 0,
      name: '',
      title: '',
      link: '',
      keywords: '',
      description: ''
    }),
    formpublished2: new Form({
      email: null
    }),
    formarchived1: new Form({
      id: 0,
      name: '',
      title: '',
      link: '',
      keywords: '',
      description: ''
    }),
    formarchived2: new Form({
      email: null
    }),
    publishedQuery: {
      key: null,
      column: 'name',
      sort: null,
      sortColumn: null,
      limit: 10,
      offset: 0,
      current: 1
    },
    archivedQuery: {
      key: null,
      column: 'name',
      sort: null,
      sortColumn: null,
      limit: 10,
      offset: 0,
      current: 1
    },
    details: {}
  }),

  computed: mapGetters({
    published_pages: 'page/published_pages',
    archived_pages: 'page/archived_pages',
    token: 'auth/token'
  }),

  methods: {
    // Confirms page publishing.
    async publishPageInit (v) {
      swal({

        title: this.$t('settings_pages__publish_page_warning_1'),
        text: this.$t('settings_pages__publish_page_warning_2') + v.name + this.$t('settings_pages__publish_page_warning_3'),
        type: 'warning',
        showCancelButton: true,
        cancelButtonColor: '#6c757d',
        confirmButtonColor: '#ffc107',
        cancelButtonText: this.$t('button__no'),
        confirmButtonText: this.$t('button__yes_publish')
      }).then((result) => {
        if (result.value) {
          this.publishPage(v)
        }
      })
    },
    // Confirms page archiving.
    async archivePageInit (v) {
      swal({
        title: this.$t('settings_pages__archive_warning_1'),
        text: this.$t('settings_pages__archive_warning_2') + v.name + this.$t('settings_pages__archive_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#343a40',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_archive'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.archivePage(v)
        }
      })
    },
    // Publishes page.
    async publishPage (v) {
      this.processed = true
      try {
        await axios.put(process.env.API_URL + '/api/page/' + v.id + '/publish', {
          headers: {Authorization: `Bearer ${this.token}`}
        })
        this.$awn.success(this.$t('settings_pages__page_published'))
        await this.$store.dispatch('page/fetchPublishedPages')
        await this.$store.dispatch('page/fetchArchivedPages')
        await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    // Archives page.
    async archivePage (v) {
      this.processed = true
      try {
        await axios.put(process.env.API_URL + '/api/page/' + v.id + '/archive', {
          headers: {Authorization: `Bearer ${this.token}`}
        })
        this.$awn.success(this.$t('settings_pages__page_archived'))
        await this.$store.dispatch('page/fetchPublishedPages')
        await this.$store.dispatch('page/fetchArchivedPages')
        await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    // Confirms page deleting.
    async deletePageInit (v) {
      swal({
        title: this.$t('settings_pages__delete_page_warning_1'),
        text: this.$t('settings_pages__delete_page_warning_2') + v.name + this.$t('settings_pages__delete_page_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_delete'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.deletePage(v)
        }
      })
    },
    // Deletes page.
    async deletePage (v) {
      this.processed = true
      try {
        await axios.delete(process.env.API_URL + '/api/page/' + v.id + '/page', {
          headers: {Authorization: `Bearer ${this.token}`}
        })

        this.$awn.success(this.$t('settings_pages__page_deleted'))
        await this.$store.dispatch('page/fetchArchivedPages')
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    async getPublishedPage (v) {
      this.$set(this, 'details', {})
      for (var i in this.published_pages) {
        if (this.published_pages[i].id === v) {
          var temp = this.template
          this.formpublished1.reset()
          _.assign(this.formpublished1, _.pick(this.published_pages[i], _.keys(temp)))
          if (this.published_pages[i]) {
            this.getPageDetail(this.published_pages[i].id)
          }
        }
      }
    },
    async getArchivedPage (v) {
      this.$set(this, 'details', {})
      for (var i in this.archived_pages) {
        if (this.archived_pages[i].id === v) {
          var temp = this.template
          this.formarchived1.reset()
          _.assign(this.formarchived1, _.pick(this.archived_pages[i], _.keys(temp)))
          if (this.archived_pages[i]) {
            this.getPageDetail(this.archived_pages[i].id)
          }
        }
      }
    },
    // Gets list of page authors.
    async getPageDetail (v) {
      try {
        const data = await this.$store.dispatch('page/getPageDetail', v)
        this.$set(this, 'details', data)
      } catch (e) {}
    },
    async addAuthorPublishedPage () {
      this.$validator.validateAll('publishedpageauthor')
        .then(async (res) => {
          if (!this.formpublished2.email || !this.regexpPattern.test(this.formpublished2.email.trim())) {
            this.$awn.alert(this.$t('settings_accounts__invalid_email_address'))
          } else {
            this.formpublished2.email = this.formpublished2.email.trim()
            if (res) {
              try {
                const { data } = await this.formpublished2.post(process.env.API_URL + '/api/page/' + this.formpublished1.id + '/author')
                this.$awn.success(this.formpublished1.name + ' - ' + this.$t(data.message))
                await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
                this.getPageDetail(this.formpublished1.id)
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
              }
            }
          }
        })
    },
    async addAuthorArchivedPage () {
      this.$validator.validateAll('archivedpageauthor')
        .then(async (res) => {
          if (!this.formarchived2.email || !this.regexpPattern.test(this.formarchived2.email.trim())) {
            this.$awn.alert(this.$t('settings_accounts__invalid_email_address'))
          } else {
            this.formarchived2.email = this.formarchived2.email.trim()
            if (res) {
              try {
                const { data } = await this.formarchived2.post(process.env.API_URL + '/api/page/' + this.formarchived1.id + '/author')
                this.$awn.success(this.formarchived1.name + ' - ' + this.$t(data.message))
                await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
                this.getPageDetail(this.formarchived1.id)
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
              }
            }
          }
        })
    },
    // Confirms author removing on published page.
    async removeAuthorInitPublishedPage (v) {
      swal({
        title: this.$t('settings_pages__remove_author_warning_1'),
        text: this.$t('settings_pages__remove_author_warning_2') + v + this.$t('settings_pages__remove_author_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_remove'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.removeAuthorPublishedPage(v)
        }
      })
    },
    // Removes author on published page.
    async removeAuthorPublishedPage (v) {
      this.processed = true
      try {
        await this.formpublished1.delete(process.env.API_URL + '/api/page/' + this.formpublished1.id + '/author/' + v, {
          headers: {Authorization: `Bearer ${this.token}`}
        })

        this.$awn.success(this.$t('settings_pages__author_removed'))
        await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
        this.getPageDetail(this.formpublished1.id)
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    // Confirms author removing on archived page.
    async removeAuthorInitArchivedPage (v) {
      swal({
        title: this.$t('settings_pages__remove_author_warning_1'),
        text: this.$t('settings_pages__remove_author_warning_2') + v + this.$t('settings_pages__remove_author_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#6c757d',
        confirmButtonText: this.$t('button__yes_remove'),
        cancelButtonText: this.$t('button__no')
      }).then((result) => {
        if (result.value) {
          this.removeAuthorArchivedPage(v)
        }
      })
    },
    // Removes author on archived page.
    async removeAuthorArchivedPage (v) {
      this.processed = true
      try {
        await this.formarchived1.delete(process.env.API_URL + '/api/page/' + this.formarchived1.id + '/author/' + v, {
          headers: {Authorization: `Bearer ${this.token}`}
        })

        this.$awn.success(this.$t('settings_pages__author_removed'))
        await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
        this.getPageDetail(this.formarchived1.id)
        this.processed = false
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.processed = false
      }
    },
    // Saves published page.
    async savePublishedPage (publishedpagesettings) {
      this.$validator.validateAll('formpublished1')
        .then(async (res) => {
          if (!this.formpublished1.name || this.formpublished1.name.trim() === '') {
            this.$awn.alert(this.$t('settings_pages__empty_page_name'))
          } else {
            this.formpublished1.name = this.formpublished1.name.trim()
            if (res) {
              try {
                const { data } = await this.formpublished1.put(process.env.API_URL + '/api/page/' + this.formpublished1.id)
                this.$awn.success(data.name + ' - ' + this.$t('edit_page__page_changes_saved'))
                await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
              }
            }
          }
        })
    },
    // Saves archived page.
    async saveArchivedPage () {
      this.$validator.validateAll('archivedpagesettings')
        .then(async (res) => {
          if (!this.formarchived1.name || this.formarchived1.name.trim() === '') {
            this.$awn.alert(this.$t('settings_pages__empty_page_name'))
          } else {
            this.formarchived1.name = this.formarchived1.name.trim()
            if (res) {
              try {
                const { data } = await this.formarchived1.put(process.env.API_URL + '/api/page/' + this.formarchived1.id)
                this.$awn.success(data.name + ' - ' + this.$t('edit_page__page_changes_saved'))
                await this.$store.dispatch('page/fetchPage', {page: this.$route.path.slice(1) || ''})
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
              }
            }
          }
        })
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
.iframe-box {
  position: relative;
  padding-top: 56.25%;
}
iframe {
  position:absolute;
  top:0;
  left:0;
  width:1366px;
  height:700px;
  -moz-transform: scale(0.35);
  -moz-transform-origin: 0 0;
  -o-transform: scale(0.35);
  -o-transform-origin: 0 0;
  -webkit-transform: scale(0.35);
  -webkit-transform-origin: 0 0;
}
</style>
