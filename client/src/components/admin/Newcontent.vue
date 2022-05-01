<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal :id="id" :title="$t('edit_page__add_existing_content')" customClass="modal fade in">
    <div class="row">
      <div class="col-md-12" id="add-content-pages">
        <DataQuery :data="contents" parent="dq-accordion-add-contents" :query="contentQuery">
          <template slot="items" slot-scope="data" v-if="data.data">
            <div class="card">
              <div class="card-header p-0">
                <div class="btn-group d-flex">
                  <v-button @click="saveContent(data.data)" class="flex-grow-0 btn btn btn-success">
                    <span class="fa fa-plus"></span>
                  </v-button>
                  <v-button class="h5 bg-transparent text-left border-0 m-0 text-dark collapsed" data-toggle="collapse" :data-target="'#addcontentpage-' + data.data.id">
                    {{ data.data.name }} {{ data.data.link ? '(' + data.data.link + ')' : '' }}
                  </v-button>
                </div>
              </div>
              <div :id="'addcontentpage-' + data.data.id" class="collapse" data-parent="#dq-accordion-add-contents">
                <div class="p-3" v-html="data.data.content"></div>
              </div>
            </div>
          </template>
        </DataQuery>
      </div>
    </div>
  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import $ from 'jquery'
export default {
  name: 'admin-newcontent',

  props: {
    id: {
      type: String,
      required: true
    },
    location: {
      type: String,
      default: 'header'
    },
    position: {
      type: Number,
      default: 0
    },
    index: {
      type: Number,
      default: 1
    }
  },

  mounted () {},

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
      {value: '_self', name: 'Same Window/Tab'},
      {value: '_blank', name: 'New Window/Tab'}
    ],
    contentQuery: {
      key: null,
      column: 'name,content',
      sort: null,
      sortColumn: null,
      limit: 5,
      offset: 0,
      current: 1
    }
  }),

  computed: mapGetters({
    contents: 'page/contents',
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
    saveContent (v) {
      let obj = {
        location: this.location,
        index: this.position,
        content: v
      }
      this.$emit('addcontent', obj, this.index)
      this.closeModal()
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
#add-content-pages .card-header .btn-group .btn.h5 {
  flex-grow: 1;
}
</style>
