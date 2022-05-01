<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="float-page-menu">
    <div class="btn-group">
      <button class="btn btn-danger trash-icon" @click="doRemove" :data-content="$t('button__remove')" data-toggle="popover" data-trigger="hover">
        <span class="fa fa-trash"></span>
      </button>
      <button class="btn btn-primary move-arrow" @mousedown="$parent.$parent.blurEditor()" :data-content="$t('button__move')" data-toggle="popover" data-trigger="hover">
        <span class="fa fa-arrows"></span>
      </button>
    </div>
  </div>
</template>

<script>
import $ from 'jquery'
export default {
  name: 'Floatpagemenu',

  props: {
    edit: {
      type: Boolean,
      default: false
    },
    close: {
      type: Boolean,
      default: false
    },
    isExtension: {
      type: Boolean,
      default: false
    }
  },

  mounted () {
    $('[data-toggle="popover"]').popover()
  },

  methods: {
    doEdit () {
      this.$emit('edit')
      this.removePopover()
    },
    doRemove () {
      this.$emit('remove')
      this.removePopover()
    },
    doClose () {
      this.$emit('close')
      this.removePopover()
    },
    removePopover () {
      // this.$el.querySelectorAll('.btn').forEach((v) => {
      //   querySelector(v).popover()
      // })
      $('[data-toggle="popover"]').popover('hide')
    }
  }
}
</script>

<style>
.float-page-menu {
  position: absolute;
  opacity: 0;
  top: 0;
  right: 0;
  z-index: 1;
}
.trash-icon{
  right: 35%;
}
.btn-group{

}
</style>
