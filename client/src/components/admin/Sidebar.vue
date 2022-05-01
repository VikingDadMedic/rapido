<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="admin-sidebar">
    <ul class="sidebar-menu">
      <template v-if="loggedIn">
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {share: true}}">
          <a data-toggle="modal" data-target="#admin-share">
            <span class="fa fa-share-alt" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__share_page') }}</span>
          </a>
        </router-link>
<!--
        <li>
          <a>
            <span class="fa fa-share-alt" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="label">Share</span>
          </a>
          <ul>
            <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {share: 'Diaspora'}}">
              <a>
                diaspora*
              </a>
            </router-link>
            <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {share: 'Mastodon'}}">
              <a>
                Mastodon
              </a>
            </router-link>
          </ul>
        </li>
 -->
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {edit: true}}" v-if="!$route.query.edit">
          <a>
            <span class="fa fa-pencil" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__edit_page') }}</span>
          </a>
        </router-link>
        <router-link tag="liquit" :to="{path: $route.path, params: {page: $route.params.page}}" v-else>
          <a>
            <span class="fa fa-times" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__close_page_edition') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {new: true}}">
          <a data-toggle="modal" data-target="#admin-newpage">
            <span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__create_new_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {duplicate: true}}">
          <a data-toggle="modal" data-target="#admin-duplicate">
            <span class="fa fa-copy" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__duplicate_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {version: true}}" v-if="pages.version.length > 1">
          <a>
            <span class="fa fa-history" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__versions') }}</span>
          </a>
        </router-link>
        <router-link tag="liarchive" :to="{path: $route.path, params: {page: $route.params.page}, query: {archive: true}}" v-if="pages.details.published === 1">
          <a>
            <span class="fa fa-moon-o" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__archive_page') }}</span>
          </a>
        </router-link>
        <router-link tag="lipublish" :to="{path: $route.path, params: {page: $route.params.page}, query: {publish: true}}" v-else>
          <a>
            <span class="fa fa-sun-o" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__publish_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {settings: 'navigation'}}">
          <a>
            <span class="fa fa-wrench" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__settings') }}</span>
          </a>
        </router-link>
        <router-link tag="liquit" :to="{path: $route.path, params: {page: $route.params.page}, query: {logout: true}}">
          <a>
            <span class="fa fa-power-off" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__logout') }}</span>
          </a>
        </router-link>
      </template>

      <template v-else>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {share: true}}">
          <a data-toggle="modal" data-target="#admin-share">
            <span class="fa fa-share-alt" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__share_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {edit: true}}" v-if="!$route.query.edit">
          <a>
            <span class="fa fa-pencil" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__edit_page') }}</span>
          </a>
        </router-link>
        <router-link tag="liquit" :to="{path: $route.path, params: {page: $route.params.page}}" v-else>
          <a>
            <span class="fa fa-times" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__close_page_edition') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {new: true}}">
          <a data-toggle="modal" data-target="#admin-newpage">
            <span class="fa fa-plus" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__create_new_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {duplicate: true}}">
          <a data-toggle="modal" data-target="#admin-duplicate">
            <span class="fa fa-copy" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__duplicate_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {version: true}}" v-if="pages.version.length > 1">
          <a>
            <span class="fa fa-history" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__versions') }}</span>
          </a>
        </router-link>
        <router-link tag="liarchive" :to="{path: $route.path, params: {page: $route.params.page}, query: {archive: true}}" v-if="pages.details.published === 1">
          <a>
            <span class="fa fa-moon-o" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__archive_page') }}</span>
          </a>
        </router-link>
        <router-link tag="lipublish" :to="{path: $route.path, params: {page: $route.params.page}, query: {publish: true}}" v-else>
          <a>
            <span class="fa fa-sun-o" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__publish_page') }}</span>
          </a>
        </router-link>
        <router-link tag="li" :to="{path: $route.path, params: {page: $route.params.page}, query: {settings: 'navigation'}}">
          <a>
            <span class="fa fa-wrench" aria-hidden="true"></span>&nbsp;&nbsp;&nbsp;<span class="label">{{ $t('sidebar__settings') }}</span>
          </a>
        </router-link>
      </template>
    </ul>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'admin-sidebar',

  computed: {
    ...mapGetters({
      pages: 'page/current_pages',
      loggedIn: 'auth/check'
    })
  }
}
</script>
