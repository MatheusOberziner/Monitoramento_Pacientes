<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar class="justify-between" style="background-color: #f7b7b7;">
          <div class="row q-px-md q-gutter-md">
            <div>
              <img src="../assets/logo.png" style="height: 80px; width: 80px;" />
            </div>
            <div class="self-center">
              <span class="text-h4 text-bold" style="color: black;">Monitoramento de Pacientes</span>
            </div>
          </div>

          <div class="row q-px-md q-gutter-md">
            <div>
              <q-btn round>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  height="1.5em"
                  viewBox="0 0 24 24"
                  color="black"
                ><path
                  fill="currentColor"
                  d="M18 10a6 6 0 0 0-12 0v8h12v-8Zm2 8.667l.4.533a.5.5 0 0 1-.4.8H4a.5.5 0 0 1-.4-.8l.4-.533V10a8 8 0 1 1 16 0v8.667ZM9.5 21h5a2.5 2.5 0 0 1-5 0Z"
                /></svg>
                <q-badge
                  floating
                  color="red"
                  rounded
                />
              </q-btn>
            </div>

            <div>
              <q-btn-dropdown round icon="person" style="color: black; height: 42px;" />
            </div>
          </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <q-page padding>
        <router-view />
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>

export default {
  name: 'MainLayout',

  setup () {
  }
}
</script>
