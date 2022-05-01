<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div :class="customClass ? customClass : 'modal fade'" :id="id" tabindex="-1" role="dialog" :aria-labelledby="id" aria-hidden="true">
    <div :class="'modal-dialog modal-dialog-centered' + (modalSize ? ' '+modalSize : '')" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <slot name="header">
            <h5 class="modal-title">
              {{ title }}
            </h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span class="fa fa-times" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__close') }}
            </button>
          </slot>
        </div>
        <div class="modal-body">
          <slot></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Modal',
  props: ['id', 'title', 'customClass', 'modalSize']
}
</script>
