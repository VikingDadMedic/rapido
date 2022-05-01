<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="admin-settings">
    <div class="admin-settings-container">
      <div class="admin-settings-container-inner h-100">
        <div class="template-wrapper h-100">
          <div class="template-header">
            <div class="container">
              <nav class="navbar fixed-top navbar-fixed-top navbar-expand-lg template-header-menu" :class="manifest.manifest ? 'navbar-default' : 'navbar-light'">
                <div class="text-center">
                  <a class="navbar-brand" style="cursor: default;"><span class="fa fa-history" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('sidebar__versions') }}</a>
                </div>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarAdmin" aria-controls="navbarAdmin" aria-expanded="false" aria-label="Toggle navigation">
                  <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarAdmin">
                  <ul class="navbar-nav ml-lg-auto d-flex align-items-center flex-row">
                    <li class="nav-item" @click="closePanel">
                      <a class="nav-link">
                        <span class="fa fa-times" aria-hidden="true"></span>&nbsp;&nbsp;{{ $t('button__close') }}
                      </a>
                    </li>
                  </ul>
                </div>
              </nav>
            </div>
          </div>

          <!-- partial -->
          <div class="template-main">
            <div class="container py-5">
              <div class="row justify-content-center">
                <div class="col-md-12" id="accordion-version">
                  <div class="card" v-for="(pg, key) in pages.version" :key="key" v-if="pages.version.length > 0">
                    <div class="card-header collapsed cursor-pointer p-0">
                      <div class="btn-group d-flex">
                        <v-button clickfn="revertPage" class="flex-grow-0 btn btn-success" @click="revertPage(pg)">
                          <h5 class="m-0 p-0">
                            {{ $t('versions__revert') }}
                          </h5>
                        </v-button>
                        <v-button class="bg-transparent text-left border-0 m-0 text-dark collapsed" data-toggle="collapse" :data-target="'#version-' + pg.id" aria-expanded="true" :aria-controls="'page-' + pg.id" clickfn="setCurrent" @click="setCurrent(pg)">
                          <h5 class="m-0 p-0">
                            {{ pg.name }}{{ $t('versions__registered on') }}
                            <span v-html="pg.created_on ? beautifyDate(pg.created_on) : ''"></span>&nbsp;&nbsp;&nbsp;
                            <span class="badge badge-danger" v-if="pg.current === 1">
                              {{ $t('versions__current_version') }}
                            </span>
                          </h5>
                        </v-button>
                      </div>
                    </div>
                    <div :id="'version-' + pg.id" class="collapse" data-parent="#accordion-version">
                      <div class="card-body">
                        <div class="row">
                          <span class="col-md-4">

                            <!-- Access Level -->
<!--                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_accounts__access_level') }}
                              </label>
                              <span class="col-md-8 col-form-label">
                                {{ pg.data.page.access_level }}
                              </span>
                            </div>
-->
                            <!-- Name -->
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_pages__name') }}
                              </label>
                              <span class="col-md-8 col-form-label">
                                {{ pg.data.page.name }}
                              </span>
                            </div>

                            <!-- Title -->
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_pages__title') }}
                              </label>
                              <span class="col-md-8 col-form-label">
                                {{ pg.data.page.title }}
                              </span>
                            </div>

                            <!-- Link -->
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_pages__link') }}
                              </label>
                              <span class="col-md-8 col-form-label">
                                {{ pg.data.page.link }}
                              </span>
                            </div>

                            <!-- Keywords -->
                            <div class="form-group row">
                              <label class="col-md-4 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_pages__keywords') }}
                              </label>
                              <span class="col-md-8 col-form-label">
                                {{ pg.data.page.keywords }}
                              </span>
                            </div>

                            <!-- Description -->
                            <div class="form-group row">
                              <label class="col-md-5 col-form-label text-md-right font-weight-bold">
                                {{ $t('settings_pages__description') }}
                              </label>
                              <span class="col-md-7 col-form-label">
                                {{ pg.data.page.description }}
                              </span>
                            </div>

                          </span>
                          <span class="col-md-8">
                            <span class="iframe-box" :ref="'ifbox-' + pg.id">
                              <iframe :src="'/' + (pg.data.page.link || '') + '?preview=true&version=true'" style="width: 1366px" :style="'transform: scale(' + ($refs['ifbox-' + pg.id].clientWidth / 1366)*10 + ')'" v-if="pg.id === current"></iframe>
                            </span>
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <h3 class="text-center" v-if="pages.version.length === 0">
                    {{ $t('versions__version_not_found') }}
                  </h3>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import $ from 'jquery'
import { mapGetters } from 'vuex'

export default {
  name: 'admin-version',

  mounted () {
    this.$nextTick(() => {
      if (this.$route.query.version) {
        $('body').css('overflow', 'hidden')
        this.setBg()
        // this.setContent(this.$route.query.version)
      }
    })
  },

  computed: {
    ...mapGetters({
      pages: 'page/current_pages',
      manifest: 'page/manifest'
    }),
    check: {
      get () {
        return this.pages.details.id ? this.pages.details.id : false
      }
    }
  },

  data: () => ({
    current: null
  }),

  beforeDestroy () {
    $('body').attr('style', '')
  },

  watch: {
    check (v) {
      if (v) {
        this.setBg()
      }
    }
  },

  methods: {
    setBg () {
      this.$nextTick(() => {
        setTimeout(() => {
          this.$el.querySelector('.admin-settings-container-inner').style.backgroundColor = $('body').css('background-color')
        }, 1000)
      })
    },
    sidebarToggle () {
      $('.admin-settings-container').toggleClass('sidebar-icon-only')
    },
    closePanel () {
      this.$router.replace({path: this.$route.path, params: {page: this.$route.params.page}})
    },
    // fetch initialization page data
    initData () {
      this.$nextTick(async () => {
        await this.$store.dispatch('page/init', {page: this.$route.path.slice(1) || ''})
      })
    },
    async revertPage (v) {
      try {
        await this.$store.dispatch('page/setPageVersion', {id: v.id})
        this.$awn.success(this.$t('versions__page_reverted_to_1') + this.pages.details.name + this.$t('versions__page_reverted_to_2') + this.beautifyDate(v.created_on) + this.$t('versions__page_reverted_to_3'))
        this.initData()
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
      }
    },
    setCurrent (v) {
      let obj = {
        details: v.data.page,
        content: v.data.page_content
      }
      this.$store.dispatch('page/setPreviewContent', obj)
      this.$set(this, 'current', v.id)
    },
    beautifyDate (v) {
      let date = new Date(v * 1000)
      let datelocale = date.toLocaleString(this.$t('locale_javascript'), {day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit'})
      return datelocale
    }
  }
}
</script>

<style scoped>
.card-header > h5 {
  cursor: pointer;
}
.card-header > .btn-group > .bg-transparent h5:after {
  font-family: 'ForkAwesome';
  content: "\f106";
  float: right;
}
.card-header > .btn-group > .collapsed.bg-transparent h5:after {
  content: "\f107";
}
.card-header > h5 {
  cursor: pointer;
}
.card-header .btn-group .btn.bg-transparent {
  flex-grow: 1;
}
.collapse.show .iframe-box {
  display: block;
}
.iframe-box {
  display: none;
  position: relative;
  padding-top: 60%;
}
iframe {
  position:absolute;
  top:0;
  left:0;
  bottom: 0;
  right: 0;
  width:1366px;
  height:700px;
  -moz-transform: scale(0.50);
  -moz-transform-origin: 0 0;
  -o-transform: scale(0.50);
  -o-transform-origin: 0 0;
  -webkit-transform: scale(0.50);
  -webkit-transform-origin: 0 0;
}
</style>
