<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>,Félix P <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="float-add-content" @mouseleave="removeShow">
    <div class="w-100">
      <div class="opt-border"></div>
      <a class="menu-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
        <span class="fa fa-plus"></span>
      </a>
      <div class="dropdown-menu">
        <a class="dropdown-item" @click.prevent="doAdd">
          <span class="fa fa-file" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('edit_page__add_new_content') }}
        </a>
        <a class="dropdown-item" data-toggle="modal" data-target="#content-modal-index" @click="doAddExisting">
          <span class="fa fa-copy" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('edit_page__add_existing_content') }}
        </a>
        <a class="dropdown-item" data-toggle="modal" data-target="#extension-modal-index" @click="doAddExtension">
          <span class="fa fa-envelope" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('edit_page__add_contact_form') }}
        </a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Floataddcontent',

  props: {
    index: {
      type: Number,
      default: 0
    },
    location: {
      type: String,
      default: 'header'
    }
  },

  methods: {
    doAdd () {
      this.$emit('add', {index: this.index, location: this.location})
    },
    doAddExisting () {
      this.$emit('addExisting', {index: this.index, location: this.location})
    },
    doAddExtension () {
      this.$emit('addExtension', {index: this.index, location: this.location})
    },
    removeShow (e) {
      let el = e.target.querySelector('.dropdown-menu')
      if (el) {
        el.className = el.className.replace('show', '')
      }
    }
  }
}
</script>

<style lang="stylus" scoped>

.float-add-content
  margin-top: 10px
  margin-bottom: 10px
  position: relative
  text-align: center
  height: 5px
  overflow: visible
  cursor: pointer
  .opt-border
    border-bottom: 1px solid #000
    position: relative
    top: 0.5em;
  .menu-toggle
    position: absolute
    top: 0.5em
    left: 50%
    color: inherit
    padding: 5px 50px
    transform: translate(-50%, -50%)
    &:hover
      color: grey
    i
      padding: 5px 10px
      background: white
  .fa-plus
    scale: 200%
    color: #007bff;
  .dropdown-menu {
    margin-top: -20px
  }
  &:hover
    opacity: 1
    transition: 0.2s all
</style>
