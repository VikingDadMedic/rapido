<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Main>
    <template v-if="(check && $route.query.edit) && !$route.query.preview">
      <Floatpagebutton @cancel="cancelEdit" @save="saveChanges" v-if="newContent.length > 0" />
      <header class="template-header">
        <navbar class="template-header-menu" />
        <div class="container">
          <div class="row" v-if="!manifest.content || manifest.content.header === 1">
            <div class="col-md-12">
              <VContent col="col-md-12" location="header" :index="1" :content="content.header[1]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
          <div class="row" v-else>
            <div :class="'col-md-' + 12 / manifest.content.header" v-for="n in manifest.content.header" :key="n">
              <VContent col="col-md-12" location="header" :index="n" :content="content.header[n]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
        </div>
      </header>
      <section class="section template-main">
        <div class="container">
          <div class="row" v-if="!manifest.content || manifest.content.main === 1 || manifest.manifest.includes('promo')">
            <div class="col-md-12">
              <VContent col="col-md-12" location="main" :index="1" :content="content.main[1]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
          <div class="row" v-if="manifest.content && manifest.content.main > 1">
            <div :class="'col-md-' + (isAbnormal() ? 12 / (manifest.content.main - 1) : 12 / manifest.content.main)" v-for="n in (isAbnormal() ? manifest.content.main - 1 : manifest.content.main)" :key="n">
              <VContent col="col-md-12" location="main" :index="isAbnormal() ? n + 1 : n" :content="content.main[(isAbnormal() ? n + 1 : n)]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
          <div class="row" v-if="manifest.content && manifest.manifest.includes('drill-down')">
            <div class="col-md-12">
              <VContent col="col-md-12" location="main" :index="1" :content="content.main[1]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
        </div>
      </section>
      <footer class="template-footer">
        <div class="container pb-5">
          <!-- <div class="row">
            <div class="col-md-12">
              <div class="template-footer-menu">
                <MenuList id="admin-menu-footer-1" className="nav" :navs="menus.footer" location="footer">
              </div>
            </div>
          </div> -->
          <div class="row" v-if="!manifest.content || manifest.content.footer === 1">
            <div class="col-md-12">
              <VContent col="col-md-12" location="footer" :index="1" :content="content.footer[1]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
          <div class="row" v-else>
            <div :class="'col-md-' + 12 / manifest.content.footer" v-for="n in manifest.content.footer" :key="n">
              <VContent col="col-md-12" location="footer" :index="n" :content="content.footer[n]" @edit="getEditor" @change="isChanged" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" :current="current" />
            </div>
          </div>
        </div>
      </footer>
      <Newcontent id="content-modal-index" :location="existing.location" :position="existing.position" :index="existing.index" @addcontent="addNewContent" v-if="user && user.access_level > 0" />
      <AddContactForm id="extension-modal-index" :location="existing.location" :position="existing.position" :index="existing.index" @addcontent="addNewContent" @addextension="addNewContent" v-if="user && user.access_level > 0" />
    </template>
    <template v-else>
      <header class="tempalte-header">
        <navbar class="template-header-menu" />
        <div class="container">
          <div class="row" v-if="!manifest.content || manifest.content.header === 1">
            <div class="col-md-12">
              <VContentStatic col="col-md-12" :content="content.header[1]" />
            </div>
          </div>
          <div class="row" v-else>
            <div :class="'col-md-' + 12 / manifest.content.header" v-for="n in manifest.content.header" :key="n">
              <VContentStatic col="col-md-12" :content="content.header[n]" />
            </div>
          </div>
        </div>
      </header>
      <section class="section template-main">
        <div class="container">
          <div class="row" v-if="!manifest.content || manifest.content.main === 1 || manifest.manifest.includes('promo')">
            <div class="col-md-12">
              <VContentStatic col="col-md-12" :content="content.main[1]" />
            </div>
          </div>
          <div class="row" v-if="manifest.content && manifest.content.main > 1">
            <div :class="'col-md-' + (isAbnormal() ? 12 / (manifest.content.main - 1) : 12 / manifest.content.main)" v-for="n in (isAbnormal() ? manifest.content.main - 1 : manifest.content.main)" :key="n">
              <VContentStatic col="col-md-12" :content="content.main[(isAbnormal() ? n + 1 : n)]" />
            </div>
          </div>
          <div class="row" v-if="manifest.content && manifest.manifest.includes('drill-down')">
            <div class="col-md-12">
              <VContentStatic col="col-md-12" :content="content.main[1]" />
            </div>
          </div>
        </div>
      </section>
      <footer class="template-footer">
        <div class="container">
          <!-- <div class="row">
            <div class="col-md-12">
              <div class="template-footer-menu">
                <MenuList id="admin-menu-footer-1" className="nav" :navs="menus.footer" location="footer">
              </div>
            </div>
          </div> -->
          <div class="row" v-if="!manifest.content || manifest.content.footer === 1">
            <div class="col-md-12">
              <VContentStatic col="col-md-12" :content="content.footer[1]" />
            </div>
          </div>
          <div class="row" v-else>
            <div :class="'col-md-' + 12 / manifest.content.footer" v-for="n in manifest.content.footer" :key="n">
              <VContentStatic col="col-md-12" :content="content.footer[n]" />
            </div>
          </div>
        </div>
      </footer>
    </template>
  </Main>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import Newcontent from '@/components/admin/Newcontent.vue'
import AddContactForm from '@/components/admin/AddContactForm.vue'
import Form from 'vform'
import axios from 'axios'
import _ from 'lodash'
import { mapGetters } from 'vuex'

export default {
  name: 'Index',

  // Set title meta information
  metaInfo () {
    let obj = {
      title: this.pages.details.title || this.pages.details.name || 'home'.replace(/\b\w/g, l => l.toUpperCase())
    }
    if (this.pages.details.theme || this.preview.details.theme) {
      let theme = this.pages.details.theme
      if (this.$route.query.preview && this.$route.query.version) {
        theme = this.preview.details.theme
      }
      if (theme) {
        let split = theme.split('/')
        let style = split[0] + '/' + split[1] + '/style/style.css'
        obj.link = [
          { rel: 'stylesheet', href: process.env.API_URL + '/extension/theme/' + style },
          { rel: 'stylesheet', href: '/static/themes.css?' + split[1] }
        ]
      }
    }
    return obj
  },

  // link: () => {
  //   let link = []
  //   if (this.pages.details.theme) {
  //     let theme = this.pages.details.theme
  //     let split = theme.split('/')
  //     let style = split[0] + '/' + split[1] + '/style/style.css'
  //     link = [
  //       { rel: 'stylesheet', href: process.env.API_URL + '/extension/' + style },
  //       { rel: 'stylesheet', href: process.env.API_URL + '/static/themes.css' }
  //     ]
  //   }
  //   return link
  // },

  components: {
    Navbar,
    Newcontent,
    AddContactForm
  },

  computed: {
    ...mapGetters({
      pages: 'page/current_pages',
      preview: 'page/preview',
      token: 'auth/token',
      user: 'auth/user',
      check: 'auth/check',
      menus: 'page/menus',
      manifest: 'page/manifest'
    }),
    contents: {
      get () {
        return this.$route.query.preview && this.$route.query.version ? this.preview.content : this.pages.content
      }
    }
  },

  watch: {
    contents (v) {
      this.initPages(v)
    },
    pages (v) {
      this.initPages(v.content)
    },
    'pages.details' () {
      this.checkManifest()
    },
    '$route' (v) {
      this.checkPermission()
    },
    current (v, old) {
      if (old) {
        this.checkContent(old)
      }
    }
  },

  data: () => ({
    current: null,
    iscontent: false,
    editorOption: {
      theme: 'snow'
    },
    content: {
      header: {
        '1': [],
        '2': [],
        '3': [],
        '4': []
      },
      main: {
        '1': [],
        '2': [],
        '3': [],
        '4': []
      },
      footer: {
        '1': [],
        '2': [],
        '3': [],
        '4': []
      }
    },
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
    newContent: [],
    existing: {
      location: null,
      position: null,
      index: null
    },
    options: {
      handle: '.move-arrow',
      disabled: false,
      group: 'content'
    }
  }),

  mounted () {
    // this.fetchContent()
    // this.initPages(this.pages)
    this.checkManifest()
  },

  methods: {
    async initPages (value) {
      let v = JSON.parse(JSON.stringify(value))
      let content = {
        header: {
          '1': [],
          '2': [],
          '3': [],
          '4': []
        },
        main: {
          '1': [],
          '2': [],
          '3': [],
          '4': []
        },
        footer: {
          '1': [],
          '2': [],
          '3': [],
          '4': []
        }
      }
      for (let i in v) {
        if (v[i].location === 'header') {
          content.header[v[i].column] = []
        } else if (v[i].location === 'main') {
          content.main[v[i].column] = []
        } else if (v[i].location === 'footer') {
          content.footer[v[i].column] = []
        }
      }
      for (let i in v) {
        if (v[i].location === 'header') {
          content.header[v[i].column].push(v[i])
        } else if (v[i].location === 'main') {
          content.main[v[i].column].push(v[i])
        } else if (v[i].location === 'footer') {
          content.footer[v[i].column].push(v[i])
        }
      }
      this.$set(this, 'content', content)
    },
    getManifest (v) {
      this.$store.dispatch('page/fetchManifest', {theme: (v || null)})
    },
    checkManifest () {
      this.$nextTick(() => {
        if (this.$route.query.preview) {
          if (this.$route.query.version) {
            this.initPages(this.preview.content)
            this.getManifest(this.preview.details.theme)
          } else {
            this.getManifest(this.pages.details.theme)
          }
        } else {
          this.getManifest(this.pages.details.theme)
        }
      })
    },
    checkContent (v) {
      for (let i in this.content) {
        for (let j in this.content[i]) {
          for (let k in this.content[i][j]) {
            this.$nextTick(() => {
              if (this.content[i][j][k]) {
                if (this.content[i][j][k].id === v) {
                  if (this.content[i][j][k].content.content && this.content[i][j][k].content.content.replace(/\s/g, '') === '') {
                    this.content[i][j].splice(k, 1)
                  }
                }
              } else {
                this.content[i][j].splice(k, 1)
              }
            })
          }
        }
      }
    },
    // async getManifest (v) {
    //   if (v.theme) {
    //     const { data } = await axios.get(process.env.API_URL + '/extension/' + v.theme, {headers: {Authorization: `Bearer ${this.token}`}})
    //     this.$set(this, 'manifest', data)
    //   } else {
    //     this.$set(this, 'manifest', {})
    //   }
    // },
    checkPermission () {
      let run = async () => {
        try {
          await axios.get(process.env.API_URL + '/api/page/' + this.pages.details.id + '/permission', {
            headers: {Authorization: `Bearer ${this.token}`}
          })
        } catch (e) {
          this.$router.replace({path: this.$route.path, query: this.$route.query})
        }
      }
      let check = setInterval(() => {
        if (this.pages.details.id) {
          run()
          clearInterval(check)
        } else {
          // console.log('blon')
        }
      }, 1000)
    },
    cancelEdit () {
      this.blurEditor()
      this.$set(this, 'newContent', [])
      this.initPages(this.contents)
    },
    getEditor (v) {
      this.current = v
    },
    blurEditor () {
      this.current = null
    },
    async saveContent () {
      try {
        await this.form.put(process.env.API_URL + '/api/content/' + this.form.id)
        this.$awn.success(this.form.name + this.$t('settings_contents__content_changes_saved'))
        this.initData()
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },
    // fetch initialization page data
    initData: _.debounce(function () {
      this.$nextTick(async () => {
        await this.$store.dispatch('page/init', {page: this.$route.path.slice(1) || ''})
      })
    }, 100),
    isChanged: _.debounce(function () {
      let result = []
      for (let i in this.content) {
        for (let j in this.content[i]) {
          for (let k in this.content[i][j]) {
            result = result.concat(this.content[i][j][k])
          }
        }
      }
      this.newContent = result
    }, 200),
    rePosition (loc, idx) {
      for (let i in this.content[loc][idx]) {
        this.content[loc][idx][i].position = parseInt(i)
      }
    },
    addNewContent (v, idx) {
      let uid = new Date()
      uid = uid.getTime()

      // Check if content is a contact-form , add json_settings to obj
      if (v.extension) {
        let obj = {
          id: uid,
          newContent: true,
          page_id: this.pages.details.id,
          location: v.location,
          position: v.index,
          column: idx,
          content: {
            name: v.content ? v.content.name : '',
            content: v.content ? v.content.content : ''
          },
          extension: v.extension ? v.extension.extension : '',
          json_settings: v.extension ? v.extension.json_settings : ''
        }
        this.content[v.location][idx].splice(v.index, 0, obj)
        this.rePosition(v.location, idx)
        this.getEditor(uid)
        this.isChanged()
      } else {
        let obj = {
          id: uid,
          newContent: true,
          page_id: this.pages.details.id,
          location: v.location,
          position: v.index,
          column: idx,
          content: {
            name: v.content ? v.content.name : '',
            content: v.content ? v.content.content : ''
          },
          extension: v.extension ? v.extension.extension : ''
        }
        this.content[v.location][idx].splice(v.index, 0, obj)
        this.rePosition(v.location, idx)
        this.getEditor(uid)
        this.isChanged()
      }
    },
    addExistingContent (v, idx) {
      this.$set(this.existing, 'location', v.location)
      this.$set(this.existing, 'position', v.index)
      this.$set(this.existing, 'index', idx)
    },
    addExtension (v, idx) {
      this.$set(this.existing, 'location', v.location)
      this.$set(this.existing, 'position', v.index)
      this.$set(this.existing, 'index', idx)
    },
    isAbnormal () {
      if (this.manifest.manifest.includes('promo') || this.manifest.manifest.includes('drill-down')) {
        return true
      } else {
        return false
      }
    },
    async saveChanges () {
      this.blurEditor()
      this.initPages(this.newContent)
      let obj = JSON.parse(JSON.stringify(this.newContent))
      for (let i in obj) {
        if (obj[i].newContent) {
          delete obj[i].id
          delete obj[i].newContent
        }
      }
      try {
        await axios.post(process.env.API_URL + '/api/content', obj)
        this.$awn.success(this.pages.details.name + ' - ' + this.$t('edit_page__page_changes_saved'))
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
      this.$set(this, 'newContent', [])
      this.initData()
    }
  }
}
</script>

<style lang="stylus">
Main {
  min-height: calc(100vh - 56px);
}
</style>
