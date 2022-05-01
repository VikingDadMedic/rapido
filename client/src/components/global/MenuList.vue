<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>,Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->

<template>
  <div>
    <draggable element="ul" :list="nav" group="menus" :options="options" class="custom-nav" :class="className" @change="moveMenu">
      <template v-for="(mn, index) in nav">
        <template v-if="!mn.sub_menu ? true : mn.sub_menu.length === 0 && !editMode">
          <li class="nav-item" :key="'url-' + mn.id" v-if="mn.link_url">
            <a class="nav-link" :href="mn.link_url.includes('//') ? mn.link_url : protocol + mn.link_url" :target="mn.link_target">
              {{ mn.link_label || mn.page.name }}
            </a>
          </li>
          <router-link tag="li" class="nav-item" :key="mn.id" :to="'/' + (mn.page.link || '')" exact v-else>
            <a class="nav-link">{{ mn.link_label || mn.page.name }}</a>
          </router-link>
        </template>
        <template v-else>
          <template v-if="!editMode || !check || !$route.query.edit">
            <router-link @mouseover.native="dropDownShow(true, index)" @mouseleave.native="dropDownShow(false, index)" tag="li" class="nav-item dropdown draggable" :key="mn.id" :to="'/' + (mn.link_url || mn.page.link || '')" exact>
              <a class="nav-link dropdown-toggle">
                {{ mn.link_label || mn.page.name }}
              </a>
              <div name= "dropdown-menu" class="dropdown-menu pt-3">
                <template v-if="editMode && (check && $route.query.edit)">
                </template>
                <template v-for="(sm, key2) in mn.sub_menu" v-else>
                  <li class="nav-item draggable2" :key="'url-' + key2" v-if="sm.link_url">
                    <a class="nav-link" :href="sm.link_url.includes('//') ? sm.link_url : protocol + sm.link_url" :target="sm.link_target">
                      {{ sm.link_label || sm.page.name }}
                    </a>
                  </li>
                  <router-link tag="li" class="nav-item draggable2" :key="key2" :to="'/' + (sm.link_url || sm.page.link || '')" exact v-else @click.prevent.native="dropDownShow(false, index)">
                    <a class="nav-link">
                      {{ sm.link_label || sm.page.name }}
                    </a>
                  </router-link>
                </template>
              </div>
            </router-link>
          </template>
          <template v-else>
            <li @mouseover="dropDownShow(true, index)" @mouseleave="dropDownShow(false, index)" class="nav-item dropdown draggable" :key="mn.id">
              <div class="nav-link input-group input-group-sm d-flex">
                <div class="input-group-prepend">
                  <button class="btn btn-danger" @click="removeParent(index)" data-trigger="hover">
                    <span class="fa fa-trash"></span>
                  </button>
                  <button class="btn btn-primary move-arrow" :data-content="$t('button__move')" data-toggle="popover" data-trigger="hover">
                    <span class="fa fa-arrows"></span>
                  </button>
                </div>
                <input type="text" class="form-control" v-model="mn.link_label" v-if="mn.link_label">
                <input type="text" class="form-control" v-model="mn.page.name" disabled v-else>
              </div>
              <draggable element="div" :list="mn.sub_menu" group="menus" :options="options2"  class="dropdown-menu pt-3" @change="moveMenu">
                <li v-if="!mn.sub_menu.length > 0">
                  <a class="nav-link text-center">
                    {{ $t('navigation_menu__no_submenu') }}
                  </a>
                </li>
                <template v-for="(sm, key2) in mn.sub_menu" v-else>
                  <li class="nav-item px-3 mb-1 draggable2" :key="'url-' + key2" v-if="sm.link_url">
                    <div class="nav-link input-group input-group-sm d-flex">
                      <div class="input-group-prepend">
                        <button class="btn btn-danger" @click="removeChild(mn.id, key2)">
                          <span class="fa fa-trash" :data-content="$t('button__remove')" data-toggle="popover" data-trigger="hover"></span>
                        </button>
                        <button class="btn btn-primary move-arrow" :data-content="$t('button__move')" data-toggle="popover" data-trigger="hover">
                          <span class="fa fa-arrows"></span>
                        </button>
                      </div>
                      <input type="text" class="form-control" v-model="sm.link_label">
                    </div>
                  </li>
                  <li class="nav-item px-3 mb-1 draggable2" :key="key2" v-else>
                    <div class="nav-link input-group input-group-sm d-flex">
                      <div class="input-group-prepend">
                        <button class="btn btn-danger" @click="removeChild(mn.id, key2)" :data-content="$t('button__remove')" data-toggle="popover" data-trigger="hover">
                          <span class="fa fa-trash"></span>
                        </button>
                        <button class="btn btn-primary move-arrow" :data-content="$t('button__move')" data-toggle="popover" data-trigger="hover">
                          <span class="fa fa-arrows"></span>
                        </button>
                      </div>
                      <input type="text" class="form-control" disabled v-model="sm.page.name">
                    </div>
                  </li>
                </template>
                <div class="btn-toolbar justify-content-center mt-2">
                  <div class="btn-group">
                    <button type="button" class="btn btn-primary btn-block" data-toggle="modal" :data-target="'#modal-' + id" @click="setMenuId(mn.id)">
                      <span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__add_submenu') }}
                    </button>
                  </div>
                </div>
              </draggable>
            </li>
          </template>
        </template>
      </template>
      <template v-if="(check && $route.query.edit) && !$route.query.preview">
        <li class="nav-item" slot="footer" v-if="!editMode">
          <a class="nav-link text-primary" @click="toggleEditMode" v-if="!editMode">
            <span class="fa fa-edit" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__edit') }}
          </a>
        </li>
        <template v-else>
          <li class="nav-item" slot="footer">
            <a class="nav-link text-primary" data-toggle="modal" :data-target="'#modal-' + id" @click="setMenuId(null)">
              <span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__add_menu') }}
            </a>
          </li>
          <li class="nav-item" slot="footer">
            <a class="nav-link text-secondary" @click="cancelEdit" v-if="editMode">
              <span class="fa fa-times" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__cancel') }}
            </a>
          </li>
          <li class="nav-item" slot="footer">
            <a class="nav-link text-success" @click="saveMenu" v-if="editMode">
              <span class="fa fa-check" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__save') }}
            </a>
          </li>
        </template>
      </template>
    </draggable>
    <Newmenu :id="'modal-' + id" :ref="id" :menuId="menuId" :subMenuId="subMenuId" :menuEdit="menuEdit" v-if="user && user.access_level > 0" @addmenu="addMenu" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import axios from 'axios'
import draggable from 'vuedraggable'
import _ from 'lodash'
import Newmenu from '@/components/admin/Newmenu.vue'
import swal from 'sweetalert2'

export default {
  name: 'MenuList',

  props: {
    id: {
      type: String,
      required: true
    },
    navs: {
      type: Object,
      default: () => ({})
    },
    location: {
      type: String,
      required: true
    },
    className: {
      type: String,
      default: 'navbar-nav'
    }
  },

  components: {
    draggable,
    Newmenu
  },

  data: () => ({
    protocol: '//',
    editMode: false,
    options: {
      handle: '.move-arrow',
      draggable: '.draggable',
      forceFallback: true,
      disabled: true
    },
    options2: {
      handle: '.move-arrow',
      draggable: '.draggable2',
      forceFallback: true,
      disabled: true
    },
    setid: {
      header: 1,
      footer: 2,
      main: 3
    },
    nav: [],
    menuId: null,
    subMenuId: null,
    menuEdit: false
  }),

  computed: {
    ...mapGetters({
      menus: 'page/menus',
      check: 'auth/check',
      user: 'auth/user'
    })
  },

  watch: {
    navs: {
      handler: 'initNav',
      deep: true
    }
  },

  methods: {
    initNav (v) {
      if (!this.editMode) {
        let obj = JSON.parse(JSON.stringify(v.item))
        for (let i in obj) {
          if (!obj[i].sub_menu) {
            obj[i].sub_menu = []
          }
        }
        this.$set(this, 'nav', obj)
      }
    },
    // fetch initialization page data
    initData: _.debounce(function () {
      this.$nextTick(async () => {
        await this.$store.dispatch('page/init', {page: this.$route.path.slice(1) || ''})
      })
    }, 100),

    setMenuId (v, sub_id = null, isEdit = false) {
      this.$set(this, 'menuId', v)
      this.$set(this, 'subMenuId', sub_id)
      this.$set(this, 'menuEdit', isEdit)
    },

    // cancelEdit() cancels the modifications
    cancelEdit () {
      this.toggleEditMode()
      this.initNav(this.navs)
      this.editMode = false
    },

    // toggleEditMode() changes editMode to TRUE
    toggleEditMode () {
      !this.editMode ? this.editMode = true : this.editMode = false
      if (this.editMode) {
        this.options.disabled = false
        this.options2.disabled = false
      } else {
        this.options.disabled = true
        this.options2.disabled = true
      }
    },

    addMenu (v) {
      let uid = new Date()
      uid = uid.getTime()
      let obj = {
        id: uid,
        newMenu: true
      }
      if (v.page) {
        obj.page_id = v.page.id
        obj.page = v.page
        obj.link_url = ''
        obj.link_label = ''
        obj.link_target = ''
      } else {
        obj.link_url = v.link.url
        obj.link_label = v.link.label
        obj.link_target = v.link.target
        obj.page = {}
      }

      if (!v.menu_id) {
        obj.position = this.nav.length
        obj.level = 1
        obj.menu_id = this.setid[this.location]
        obj.sub_menu = []
        this.nav.push(obj)
      } else {
        let index = _.findIndex(this.nav, (o) => { return o.id === parseInt(v.menu_id) })
        let index2 = null
        if (v.submenu_id) {
          index2 = _.findIndex(this.nav[index].sub_menu, (o) => { return o.id === parseInt(v.submenu_id) })
        }
        if (!v.isEdit) {
          obj.level = 2
          obj.menu_id = this.setid[this.location]
          if (this.nav[index].sub_menu) {
            obj.position = this.nav[index].sub_menu.length
            this.nav[index].sub_menu.push(obj)
          } else {
            obj.position = 0
            this.nav[index].sub_menu = []
            this.nav[index].sub_menu.push(obj)
          }
        } else {
          if (v.submenu_id) {
            _.assignIn(this.nav[index].sub_menu[index2], obj)
          } else {
            _.assignIn(this.nav[index], obj)
          }
        }
      }
    },

    moveMenu (v) {
      let obj = {}
      let dumEl = v.added ? v.added.element : (v.moved ? v.moved.element : v.removed.element)
      let dumPos = v.added ? v.added.newIndex : (v.moved ? v.moved.newIndex : v.removed.newIndex)
      if (v.added || v.moved) {
        obj.level = 1
        obj.id = dumEl.id
        obj.page_id = dumEl.page_id
        obj.position = dumPos
        for (let i in this.nav) {
          if (this.nav[i].id === obj.id) {
            _.merge(this.nav[i], obj)
          }
          this.nav[i].position = parseInt(i)
        }
      } else if (v.removed) {}
    },

    // removeChild takes id,key and splices it if logged in user is creator / admin
    removeChild (id, key) {
      let index = _.findIndex(this.nav, (o) => { return o.id === parseInt(id) })
      if (this.nav[index].sub_menu[0].user_id === this.user.id) {
        try {
          this.nav[index].sub_menu.splice(key, 1)
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
      if (this.user.access_level >= 10) {
        try {
          this.nav[index].sub_menu.splice(key, 1)
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      } if (this.nav[index].sub_menu[0].user_id !== this.user.id) {
        this.$awn.alert(this.$t('navigation_menu__cannot_edit_menu'))
      }
    },

    // RemoveParent takes menu_item index and deletes it if logged in user is creator / admin
    removeParent (index) {
      const del = () => {
        try {
          this.nav.splice(index, 1)
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
      // Begin condition if menu_item has 1 or more submenus
      if (this.nav[index].sub_menu.length > 0) {
        //  If logged in user is item creator , display swal and delete
        if (this.nav[index].user_id === this.user.id) {
          swal({
            title: this.$t('navigation_menu__remove_menu_warning_1'),
            text: this.$t('navigation_menu__remove_menu_warning_2') + this.nav[index].page.name + this.$t('navigation_menu__remove_menu_warning_3') + this.$t('navigation_menu__remove_menu_warning_4'),
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#dc3545',
            cancelButtonColor: '#6c757d',
            confirmButtonText: this.$t('button__yes_remove'),
            cancelButtonText: this.$t('button__no')
          }).then((result) => {
            if (result.value) {
              del()
            }
          })
        //  If logged in user is Admin, display swal and delete
        } if (this.user.access_level >= 10) {
          swal({
            title: this.$t('navigation_menu__remove_menu_warning_1'),
            text: this.$t('navigation_menu__remove_menu_warning_2') + this.nav[index].page.name + this.$t('navigation_menu__remove_menu_warning_3') + this.$t('navigation_menu__remove_menu_warning_4'),
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#dc3545',
            cancelButtonColor: '#6c757d',
            confirmButtonText: this.$t('button__yes_remove'),
            cancelButtonText: this.$t('button__no')
          }).then((result) => {
            if (result.value) {
              del()
            }
          })
        // Display error, user not item creator nor admin
        } else if (this.user.access_level < 10 && this.user.id !== this.nav[index].user_id) {
          this.$awn.alert(this.$t('navigation_menu__cannot_edit_menu'))
        }// eslint-disable-next-line
      }
      //  Begin Condition for single menu_item
      else {
        //  If logged in user is creator , display swal as normal
        if (this.nav[index].user_id === this.user.id) {
          swal({
            title: this.$t('navigation_menu__remove_menu_warning_1'),
            text: this.$t('navigation_menu__remove_menu_warning_2') + this.nav[index].page.name + this.$t('navigation_menu__remove_menu_warning_4'),
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#dc3545',
            cancelButtonColor: '#6c757d',
            confirmButtonText: this.$t('button__yes_remove'),
            cancelButtonText: this.$t('button__no')
          }).then((result) => {
            if (result.value) {
              del()
            }
          })
        //  If logged in user is Admin, display swal
        } if (this.user.access_level >= 10) {
          swal({
            title: this.$t('navigation_menu__remove_menu_warning_1'),
            text: this.$t('navigation_menu__remove_menu_warning_2') + this.nav[index].page.name + this.$t('navigation_menu__remove_menu_warning_4'),
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#dc3545',
            cancelButtonColor: '#6c757d',
            confirmButtonText: this.$t('button__yes_remove'),
            cancelButtonText: this.$t('button__no')
          }).then((result) => {
            if (result.value) {
              del()
            }
          })
          //  Display error, user not creator nor admin
        } else if (this.user.access_level < 10 && this.user.id !== this.nav[index].user_id) {
          this.$awn.alert(this.$t('navigation_menu__cannot_edit_menu'))
        }
      } //  End condition for single menu_item
    },

    async saveMenu () {
      let obj = JSON.parse(JSON.stringify(this.nav))
      for (let i in obj) {
        if (obj[i].newMenu) {
          delete obj[i].id
          delete obj[i].newMenu
        }
        for (let o in obj[i].sub_menu) {
          if (obj[i].sub_menu[o].newMenu) {
            delete obj[i].sub_menu[o].id
            delete obj[i].sub_menu[o].newMenu
          }
        }
      }
      let finalNav = JSON.parse(JSON.stringify(this.menus))
      if (!finalNav[this.location].item && !finalNav[this.location].id) {
        finalNav[this.location] = {
          access_level: null,
          id: this.setid[this.location],
          item: [],
          name: this.location
        }
      }
      finalNav[this.location].item = obj
      let finalArr = []
      for (let i in finalNav) {
        finalArr.push(finalNav[i])
      }
      try {
        const { data } = await axios.post(process.env.API_URL + '/api/menu', finalArr)
        this.$awn.success(this.$t(data.message))
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
      this.toggleEditMode()
      this.initData()
      this.initNav(this.navs)
    },

    //  Adds "show" attribute to the submenu class if bool == true
    dropDownShow (bool, i) {
      if (bool) {
        if (this.editMode) {
          for (let menu of document.querySelectorAll('.nav-item.dropdown.draggable')) {
            menu.children[1].classList.add('show')
          }
        }
        if (!this.editMode) {
          for (let menu of document.querySelectorAll('.nav-item.dropdown.draggable')) {
            menu.children[1].classList.remove('show')
            menu.onmouseover = function () {
              menu.children[1].classList.add('show')
            }
            menu.onmouseleave = function () {
              menu.children[1].classList.remove('show')
            }
          }
        }
      }
    },
    async fetchTheme () {
      const { data } = await axios.get(process.env.API_URL + '/api/theme', {headers: {Authorization: `Bearer ${this.token}`}})
      console.log(data)
    }
  },

  mounted () {
    this.fetchTheme()
  },

  updated () {
    // Cancel the menu edit if page editing closed
    if (this.$route.query.edit === undefined && this.editMode) {
      this.cancelEdit()
    }
  }
}
</script>

<style lang="stylus">
.custom-nav .dropdown
  > .dropdown-menu
    margin-top: 0
  // .dropdown-toggle:active
  //   pointer-events: none
.show
  display: block
.dropdown-menu.pt-3.show
  width: auto

.fa-arrows
  color: white

.navbar-light .navbar-nav .active > .nav-link, .navbar-light .navbar-nav .nav-link.active, .navbar-light .navbar-nav .nav-link.show, .navbar-light .navbar-nav .show > .nav-link
  color: #007bff;

</style>
