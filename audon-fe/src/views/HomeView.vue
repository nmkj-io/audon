<script>
import { useMastodonStore } from "../stores/mastodon";
import axios from "axios";

export default {
  setup() {
    return {
      donStore: useMastodonStore(),
    };
  },
  data() {
    return {
      query: "",
    };
  },
  methods: {
    async onLogout() {
      // if (!confirm(this.$t("logoutConfirm"))) return;

      try {
        const resp = await axios.post("/app/logout");
        if (resp.status === 200) {
          this.donStore.$reset();
          this.$router.push({ name: "login" });
        }
      } catch (error) {
        console.log(error);
      } finally {
        this.donStore.$reset();
        this.$router.push({ name: "login" });
      }
    },
  },
};
</script>

<template>
  <main>
    <div class="text-right">
      <v-btn variant="outlined" color="red" @click="onLogout">
        {{ $t("logout") }}
      </v-btn>
    </div>
    <div class="text-center my-10">
      <v-avatar class="rounded" size="100">
        <v-img
          :src="donStore.userinfo?.avatar"
          :alt="donStore.userinfo?.displayName"
        >
        </v-img>
      </v-avatar>
      <h2 class="mt-5">
        {{ donStore.userinfo?.displayName }}
      </h2>
      <div>
        <a :href="donStore.userinfo?.url" class="plain">{{
          donStore.myWebfinger
        }}</a>
      </div>
    </div>
    <v-row class="text-center" justify="center">
      <!-- <v-col cols="12">
        <v-text-field v-mode="query"></v-text-field>
      </v-col> -->
      <v-col cols="12">
        <v-btn block :to="{ name: 'create' }" color="indigo">{{
          $t("createNewRoom")
        }}</v-btn>
      </v-col>
    </v-row>
  </main>
</template>
