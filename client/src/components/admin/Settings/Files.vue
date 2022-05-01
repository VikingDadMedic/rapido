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
      <div class="card">
        <div class="card-body">

          <!-- Upload file(s) -->
          <div v-if="!asModal">
            <form @submit.prevent="saveImages" @keydown="form.onKeydown($event)">
              <div class="form-group row">
                <span class="col-md-12 d-flex justify-content-end">
                  <v-button nativeType="button" :loading="form.busy" clickfn="" @click="$refs.fileinput.click()" class="btn btn-primary">
                    <span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('settings_files__upload_file') }}
                  </v-button>
                  <input type="file" ref="fileinput" name="files" class="form-control upload-image d-none" @change="onFileChanged" multiple>
                </span>
              </div>
            </form>
          </div>
          <hr>

          <!-- Filter files -->
          <div class="form-group row">
            <span class="col-md-auto">
              <span class="input-group mb-6">
                <span class="input-group-prepend">
                  <span class="input-group-text">
                    <input type="checkbox" name="show" v-model="show.unused" value="unused" id="unused-checkbox">
                  </span>
                </span>
                <label class="form-control d-inline-flex w-auto px-3" for="unused-checkbox">
                  {{ $t('settings_files__unused_files') }}
                </label>
                <span class="input-group-prepend">
                  <span class="input-group-text">
                    <input type="checkbox" name="show" v-model="show.used" value="used" id="used-checkbox">
                  </span>
                </span>
                <label class="form-control d-inline-flex w-auto px-3" for="used-checkbox">
                  {{ $t('settings_files__files_in_use') }}
                </label>
              </span>
            </span>
          </div>

          <div class="form-group row">
            <table class="table">
              <thead>
                <tr>
                  <th></th>
                  <th></th>
                  <th>
                    {{ $t('settings_files__file_name') }}
                  </th>
                  <th v-if="!asModal">
                    {{ $t('settings_files__file_type') }}
                  </th>
                  <th></th>
                </tr>
              </thead>
              <DataQueryTable :data="images" parent="dq-accordion-media" :query="mediaQuery" :columnspan="5" v-if="$route.query.settings === 'files' || asModal">
                <template slot="items" slot-scope="data" v-if="data.data">
                  <td>
                    <div class="thumbnail float-left">
                      <template v-if="data.data.type == 'image'">
                        <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" target="_blank" :title="$t('settings_files__show_in_new_tab')"><img class="card-img-top" :src="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" :alt="data.data.name"></a>
                      </template>
                      <template v-else-if="data.data.type == 'audio'">
                        <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" target="_blank" :title="$t('settings_files__show_in_new_tab')" class="d-block text-center">
                          <span class="fa fa-music" aria-hidden="true"></span>
                        </a>
                      </template>
                      <template v-else-if="data.data.type == 'video'">
                        <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" target="_blank" :title="$t('settings_files__show_in_new_tab')" class="d-block text-center">
                          <span class="fa fa-video" aria-hidden="true"></span>
                        </a>
                      </template>
                      <template v-else-if="data.data.type == 'document'">
                        <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" target="_blank" :title="$t('settings_files__show_in_new_tab')" class="d-block text-center">
                          <span class="fa fa-file" aria-hidden="true"></span>
                        </a>
                      </template>
                      <template v-else>
                        <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" target="_blank" :title="$t('settings_files__show_in_new_tab')" class="d-block text-center">
                          <span class="fa fa-file" aria-hidden="true"></span>
                        </a>
                      </template>
                    </div>
                  </td>
                  <td>
                    <div class="btn-group" role="group" aria-label="Basic example" v-if="!asModal">
                      <a class="btn btn-flat" :title="$t('settings_files__copy_URL_to_clipboard')" @click="doCopy(data.data, data.keys)">
                        <span class="fa" :class="copied !== data.keys ? 'fa-copy' : 'fa-check'"></span>
                      </a>
                    </div>
                    <div class="btn-group" role="group" v-else>
                      <a class="btn btn-flat" title="choose" @click="chooseImage(data.data)">
                        <span class="fa fa-check"></span>
                      </a>
                    </div>
                  </td>
                  <td>
                    <a :href="data.data.url.includes('//') ? data.data_url : protocol +  data.data.url" class="text-dark" target="_blank" :title="$t('settings_files__show_in_new_tab')">
                      {{ data.data.name }}
                    </a>
                  </td>
                  <td v-if="!asModal">
                    <template v-if="data.data.type == 'image'">
                      {{ $t('settings_files__type_image') }}
                    </template>
                    <template v-else-if="data.data.type == 'audio'">
                      {{ $t('settings_files__type_audio') }}
                    </template>
                    <template v-else-if="data.data.type == 'video'">
                      {{ $t('settings_files__type_video') }}
                    </template>
                    <template v-else-if="data.data.type == 'document'">
                      {{ $t('settings_files__type_office_document') }}
                    </template>
                    <template v-else>
                      {{ $t('settings_files__type_other') }}
                    </template>
                  </td>
                  <td>
                    <div class="btn-group" role="group" aria-label="Basic example" v-if="!asModal">
                      <a class="btn btn-flat" :title="$t('button__delete')" @click="deleteImage(data.data)">
                        <span class="fa fa-trash"></span>
                      </a>
                    </div>
                  </td>
                </template>
              </DataQueryTable>
            </table>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import axios from 'axios'
import swal from 'sweetalert2'
import _ from 'lodash'
import { EventBus } from '@/plugins/event-bus.js'
export default {
  name: 'admin-settings-gallery',

  props: {
    asModal: {
      type: Boolean,
      default: false
    }
  },

  data: () => ({
    protocol: '//',
    template: {
      id: null,
      files: null
    },
    form: new Form({
      id: null,
      files: null
    }),
    files: new FormData(),
    copied: null,
    show: {
      used: true,
      unused: true
    },
    mediaQuery: {
      key: null,
      column: 'name,type',
      sort: null,
      sortColumn: null,
      limit: 10,
      offset: 0,
      current: 1
    }
  }),

  computed: mapGetters({
    images: 'gallery/images',
    token: 'auth/token'
  }),

  watch: {
    'show.used' (v) {
      if (this.show.unused === false && this.show.used === false) {
        this.show.unused = true
      }
      this.fetchAllImage()
    },
    'show.unused' (v) {
      if (this.show.unused === false && this.show.used === false) {
        this.show.used = true
      }
      this.fetchAllImage()
    }
  },

  mounted () {
    if (this.token) {
      this.fetchAllImage()
    }
  },

  methods: {
    onFileChanged (e) {
      this.files = new FormData()
      for (let i in e.target.files) {
        if (e.target.files[i].type) {
          let file = e.target.files[i]
          this.files.append('files', file)
        }
      }
      if (e.target.files.length > 0) {
        this.saveImages()
      }
    },
    deleteImage (v) {
      swal({
        title: this.$t('settings_files__delete_file_warning_1'),
        text: this.$t('settings_files__delete_file_warning_2') + v.name + this.$t('settings_files__delete_file_warning_3'),
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#dc3545',
        cancelButtonColor: '#d6c757d',
        confirmButtonText: this.$t('button__yes_delete'),
        cancelButtonText: this.$t('button__no')
      }).then(async (result) => {
        if (result.value) {
          let obj = {
            name: v.name,
            type: v.type
          }
          try {
            await axios.delete(process.env.API_URL + '/api/file/media', {params: obj})
            this.$awn.success(this.$t('settings_files__file_deleted'))
            this.fetchAllImage()
          } catch (e) {
            this.$awn.alert(this.$t(e.response.data.message))
          }
        }
      })
    },
    chooseImage (data) {
      EventBus.$emit('my-event', data)
      this.$parent.$parent.closeModal()
    },
    saveImages () {
      this.$validator.validateAll()
        .then(async (res) => {
          if (res) {
            try {
              const { data } = await axios.post(process.env.API_URL + '/api/file/media', this.files, { headers: {'Content-Type': 'application/x-www-form-urlencoded'} })
              this.$awn.success(this.$t(data.message))
              this.fetchAllImage()
            } catch (e) {
              this.$awn.alert(this.$t(e.response.data.message))
            }
          }
        })
    },
    // debounce() prevent to refresh the display if we click twice on a checkbox in less than x milliseconds
    fetchAllImage: _.debounce(async function () {
      await this.$store.dispatch('gallery/getImageList', {show: this.show.used ? (this.show.unused ? 'all' : 'used') : (this.show.unused ? 'unused' : 'none')})
    }, 300),
    doCopy (v, key) {
      this.$copyText(v.url).then((e) => {
        this.successCopy(key)
      })
    },
    successCopy (v) {
      this.copied = v
    }
  }
}
</script>

<style lang="stylus" scoped>
.thumbnail
  width: 50px
  padding: 5px
  background: #ecf0f1
  img
    height: auto
    max-width: 100%
.w-auto
  width: auto
</style>
