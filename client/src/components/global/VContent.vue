<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div>
    <draggable element="div" :list="content" :options="options" :class="{'row page-container py-3': content.length > 0, 'row page-container-blank': content.length === 0}" @change="changeStructure($event, location)" v-if="content.length > 0">
      <div :class="col + ' draggable'" v-for="(pg, key) in content" :key="pg.id">
        <Floataddcontent :index="0" :location="location" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" v-if="key === 0" />
        <div class="ql-editor page-content">
          <Floatpagemenu @edit="getEditor(pg.id)" @remove="removeContent(key, 'location')" @close="blurEditor" :edit="current === pg.id" :isExtension="pg.extension ? true : false" :close="current === pg.id" />
          <template v-if="!pg.extension">
            <input type="text" class="form-control" v-model="pg.content.name" @change="isChanged" v-if="current === pg.id">
            <quill-editor :ref="'editor-' + pg.id" :options="editorOption" v-model="pg.content.content" @change="isChanged" v-if="current === pg.id"></quill-editor>
            <div v-html="pg.content.content" @click="getEditor(pg.id)" v-if="current !== pg.id" :style="!pg.content.content || pg.content.content === '' ?'height: 30px' : ''"></div>
          </template>
          <template v-else>
            <ContactForm :to="pg.json_settings[0].value" />
          </template>
        </div>
        <Floataddcontent :index="parseInt(key)+1" :location="location" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" />
      </div>
    </draggable>
    <VContentEmpty :col="col" :location="location" :index="index" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" v-if="content.length === 0" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import AsyncComponent from './AsyncComponent.vue'
import draggable from 'vuedraggable'
import axios from 'axios'
import _ from 'lodash'

export default {
  name: 'VContent',

  metaInfo () {
    let obj = []
    for (let i in this.content) {
      if (this.content[i].extension) {
        obj.push({ src: '/extensions/' + this.content[i].extension + '/dist/umd/' + this.content[i].extension + '.min.js', type: 'application/html' })
      }
    }
    return {
      script: obj
    }
  },

  computed: {
    ...mapGetters({
      token: 'auth/token',
      contactforms: 'page/contactforms',
      user: 'auth/user'
    })
  },
  props: {
    content: {
      type: Array,
      default: () => ([])
    },
    location: {
      type: String,
      required: true
    },
    index: {
      type: Number,
      required: true
    },
    col: {
      type: String,
      default: 'col-md-12'
    },
    current: {
      default: null
    }
  },

  components: {
    draggable,
    AsyncComponent
  },

  data: () => ({
    editorOption: {
      theme: 'snow'
    },
    options: {
      handle: '.move-arrow',
      disabled: false,
      forceFallback: true,
      group: 'content'
    },
    isLoaded: {},
    api_url: process.env.API_URL,
    autoselect: true
  }),

  watch: {
    content (v) {
      // this.initCheckComponent(v)
    }
  },

  mounted () {
    this.startEdit()
    this.clickFirst()
    this.autoSelect()
    this.initCheckComponent(this.content)
  },

  updated () {
    this.startEdit()
    this.clickFirst()
    if (this.autoselect) {
      this.autoSelect()
    }
  },

  methods: {

    // Start editing
    startEdit () {
      const editor = document.querySelector('.col-md-12.draggable:first-child .ql-container.ql-snow .ql-editor')
      if (editor !== null) {
        editor.focus()
        document.querySelector('html').onclick = () => {
          this.startEdit = () => {}
        }
      }
    },
    clickFirst () {
      const firstContent = document.querySelector('.col-md-12.draggable:first-child .ql-editor div:nth-child(2)')
      if (firstContent !== null) {
        firstContent.click()
        document.querySelector('html').onclick = () => {
          this.clickFirst = () => {}
        }
      }
    },
    autoSelect () {
      let openEditor = '.col-md-12.draggable .ql-container.ql-snow .ql-editor'
      if (document.querySelector(openEditor) !== null) {
        for (let block of document.querySelectorAll(openEditor)) {
          block.focus()
          block.parentNode.parentNode.parentNode.childNodes[2].onclick = (e) => {
            this.autoselect = false
            block.onclick = () => {
              this.autoselect = true
            }
          }
        }
      }
    },

    // Change Page structure
    changeStructure (v, loc) {
      let obj = {}
      let dumEl = v.added ? v.added.element : (v.moved ? v.moved.element : v.removed.element)
      let dumPos = v.added ? v.added.newIndex : (v.moved ? v.moved.newIndex : v.removed.newIndex)
      if (v.added || v.moved) {
        obj.id = dumEl.id
        obj.page_id = dumEl.page_id
        obj.location = loc
        obj.position = dumPos
        for (let i in this.content) {
          if (this.content[i].id === obj.id) {
            _.merge(this.content[i], obj)
          }
          this.content[i].position = parseInt(i)
        }
        this.isChanged()
      } else if (v.removed) {}
    },
    getEditor (v) {
      this.$emit('edit', v)
    },
    async getView (v, idx) {
      let data = await axios.get(process.env.API_URL + '/client/src/extensions/' + v.extension + '/src/' + v.extension + '.vue')
      this.$set(this.content[idx], 'extensioncontent', data.data.replace(/<\?php echo RAZOR_BASE_URL \?>/g, process.env.API_URL + '/').replace(/<\?[=|php]?[\s\S]*?\?>/g, ''))
      return true
    },
    initCheckComponent (v) {
      for (let i in v) {
        if (v[i].extension) {
          // this.checkComponent(v[i], i)
        }
      }
    },
    // Checks Component
    /* checkComponent (v, idx) {
      if (v.extension) {
        let check = setInterval(() => {
          if (window[this.$toTitleCase(this.$toCamelCase(v.extension.split('/')[2]))]) {
            this.$set(this.isLoaded, idx, true)
            clearInterval(check)
          }
        }, 500)
      }
    }, */
    rePosition () {
      for (let i in this.content) {
        this.content[i].position = parseInt(i)
      }
    },
    removeContent (index) {
      this.content.splice(index, 1)
      this.rePosition()
      this.isChanged()
    },
    isChanged () {
      this.$emit('change', true)
    },
    addNewContent (v) {
      this.$emit('add', v, this.index)
    },
    addExistingContent (v) {
      this.$emit('addExisting', v, this.index)
    },
    addExtension (v) {
      this.$emit('addExtension', v, this.index)
    },
    blurEditor () {
      // console.log('blur')
      this.$emit('edit', null)
    }
  }
}

</script>

<style lang="stylus">
.page-content
  position: relative
  &:hover > .float-page-menu
    opacity: 0.7
    transition: 0.2s all
  // &:active {
  //   -webkit-box-shadow: 0px 0px 5px 0px rgba(144, 144, 144, 0.39);
  //   box-shadow: 0px 0px 5px 0px rgba(144, 144, 144, 0.39);
  // }

.page-container
  min-height: 100px
  // -webkit-box-shadow: 0px 0px 5px 0px rgba(144, 144, 144, 0.5)
  // box-shadow: 0px 0px 5px 0px rgba(144, 144, 144, 05)

.page-container-blank
  min-height: 10px

// .ql-container.ql-snow, .ql-toolbar.ql-snow
//   border: none


.template-wrapper main>section.section.template-main, header .float-add-content
  display: none
</style>
