<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="float-page-button">
    <div class="btn-group">
      <button class="btn btn-success" @click="doSave">
        <span class="fa fa-check" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('edit_page__save_changes') }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Floatpagebutton',

  methods: {
    doSave () {
      this.$emit('save')
    }
    // doCancel () {
    //   this.$emit('cancel')
    // }
  }
}
</script>

<style>
.float-page-button {
  position: fixed;
  /* opacity: 0; */
  bottom: 10px;
  right: 0;
  z-index: 5;
}
</style>
