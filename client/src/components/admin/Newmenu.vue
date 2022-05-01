<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal :id="id" :title="$t('navigation_menu__add_menu_or_submenu')" customClass="modal fade in">
    <div class="row">
      <div class="col-md-12">
        <ul class="nav nav-tabs">
          <li class="nav-item">
            <a class="nav-link active" data-toggle="tab" href="#add-menu-pages" role="tab" @click="currentTab('pages')">
              {{ $t('settings_pages') }}
            </a>
          </li>
          <li class="nav-item">
            <a class="nav-link" data-toggle="tab" href="#add-menu-url" role="tab" @click="currentTab('url')">
              {{ $t('navigation_menu__url') }}
            </a>
          </li>
        </ul>
        <div class="row tab-content">

          <!-- Tab Add page in navigation menu -->
          <div class="col-md-12 tab-pane fade show active" id="add-menu-pages">
            <DataQuery :data="pages" parent="dq-accordion-add-menu" :query="pageQuery">
              <template slot="items" slot-scope="data" v-if="data.data">
                <div class="card">
                  <div class="card-header p-0">
                    <div class="btn-group d-flex">
                      <v-button clickfn="saveMenu" @click="saveMenu(data.data)" class="flex-grow-0 btn btn-success">
                        <span class="fa fa-plus"></span>
                      </v-button>
                      <v-button class="h5 bg-transparent text-left border-0 m-0 text-dark collapsed" data-toggle="collapse" :data-target="'#addmenupage-' + data.data.id" clickfn="getMenu" @click="getMenu(data.data.id)">
                        {{ data.data.name }} {{ data.data.link ? '(' + data.data.link + ')' : '' }}
                      </v-button>
                    </div>
                  </div>
                  <div :id="'addmenupage-' + data.data.id" class="collapse" data-parent="#dq-accordion-add-menu">
                    <div class="iframe-box" :ref="'ifbox-' + data.data.id">
                      <iframe :src="'/' + (data.data.link || '') + '?preview=true'" style="width: 1366px" v-if="selectedMenu === data.data.id" :style="'transform: scale(' + ($refs['ifbox-' + data.data.id].clientWidth / 1366) > 0 ?($refs['ifbox-' + data.data.id].clientWidth / 1366) : 0.34+ ')'"></iframe>
                    </div>
                  </div>
                </div>
              </template>
            </DataQuery>
          </div>

          <!-- Tab Add URL in navigation menu -->
          <div class="col-md-12 tab-pane fade" id="add-menu-url">
            <div class="clearfix"></div>
            <br>

            <!-- URL -->
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('navigation_menu__url') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.url" type="text" name="url" class="form-control">
              </span>
            </div>

            <!-- Label -->
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('navigation_menu__label') }}
              </label>
              <span class="col-md-8">
                <input v-model="form.label" type="text" name="label" class="form-control">
              </span>
            </div>

            <!-- Target -->
            <div class="form-group row">
              <label class="col-md-4 col-form-label text-md-right">
                {{ $t('navigation_menu__target') }}
              </label>
              <span class="col-md-8">
                <select v-model="form.target" type="text" name="target" class="form-control">
                  <option v-for="(op, key) in target" :value="op.value" :key="key">
                    {{ $t(op.name) }}
                  </option>
                </select>
              </span>
            </div>

            <!-- Add URL button -->
            <div class="form-group row">
              <span class="col-md-4 offset-md-4">
                <v-button :loading="form.busy" clickfn="addURL" slot="footer" @click="addURL" v-if="current === 'url'" class="btn btn-success">
                  {{ $t('button__add_url') }}
                </v-button>
              </span>
            </div>

          </div>
        </div>
      </div>
    </div>
  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import $ from 'jquery'
export default {
  name: 'admin-newmenu',

  props: {
    id: {
      type: String,
      required: true
    },
    menuId: {
      type: Number,
      default: null
    },
    subMenuId: {
      type: Number,
      default: null
    },
    menuEdit: {
      type: Boolean,
      default: false
    }
  },

  data: () => ({
    current: 'pages',
    template: {
      url: '',
      label: '',
      target: '_self'
    },
    selectedMenu: null,
    form: new Form({
      url: '',
      label: '',
      target: '_self'
    }),
    target: [
      {value: '_self', name: 'navigation_menu__same_window_tab'},
      {value: '_blank', name: 'navigation_menu__new_window_tab'}
    ],
    pageQuery: {
      key: null,
      column: 'name',
      sort: null,
      sortColumn: null,
      limit: 5,
      offset: 0,
      current: 1
    }
  }),

  computed: mapGetters({
    pages: 'page/pages',
    access_level: 'page/access_level',
    token: 'auth/token'
  }),

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {
    currentTab (v) {
      this.$set(this, 'current', v)
    },
    getMenu (v) {
      this.selectedMenu = v
    },
    // Save new page
    async saveMenu (v) {
      let obj = {
        menu_id: this.menuId,
        submenu_id: this.subMenuId,
        page: v,
        isEdit: this.menuEdit
      }
      this.$emit('addmenu', obj)
      this.closeModal()
    },
    addURL () {
      if (!this.form.url || this.form.url.trim() === '') {
        this.$awn.alert(this.$t('navigation_menu__empty_URL'))
      } else if (!this.form.label || this.form.label.trim() === '') {
        this.$awn.alert(this.$t('navigation_menu__empty_label'))
      } else {
        this.form.url = this.form.url.trim()
        this.form.label = this.form.label.trim()
        let obj = {
          menu_id: this.menuId,
          submenu_id: this.subMenuId,
          link: {
            url: this.form.url,
            label: this.form.label,
            target: this.form.target
          },
          isEdit: this.menuEdit
        }
        this.$emit('addmenu', obj)
        this.closeModal()
      }
    },
    // launch modal
    launchModal () {
      $('#' + this.id).modal('show')
    },
    // close modal
    closeModal () {
      $('#' + this.id).modal('hide')
    },
    // destroy modal
    destroyModal () {
      $('#' + this.id).modal('dispose')
    }
  }
}
</script>

<style scoped>
.card-header > h5 {
  cursor: pointer;
}
.card-header > .btn-group > .h5:after {
  font-family: 'ForkAwesome';
  content: "\f106";
  float: right;
}
.card-header > .btn-group > .collapsed.h5:after {
  content: "\f107";
}
.collapse.show .iframe-box {
  display: block;
}
.iframe-box {
  display: none;
  position: relative;
  padding-top: 50%;
  height: 100%;
  width: 100%;
}
iframe {
  position:absolute;
  top:0;
  left:0;
  bottom: 0;
  right: 0;
  width:1366px;
  height:700px;
  -moz-transform: scale(0.34);
  -moz-transform-origin: 0 0;
  -o-transform: scale(0.34);
  -o-transform-origin: 0 0;
  -webkit-transform: scale(0.34);
  -webkit-transform-origin: 0 0;
}
#add-menu-pages .card-header .btn-group .btn.h5 {
  flex-grow: 1;
}
</style>
