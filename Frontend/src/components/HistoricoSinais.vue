<template>
  <q-dialog ref="dialogRef" @hide="onDialogHide">
    <div
      v-if="dadosCarregados"
      class="col-12 col-md-12"
      style="min-width: 50% !important;"
    >
      <q-card>
        <q-card-section class="q-px-lg">
          <div class="row justify-between">
            <div class="text-primary text-h5">
              Histórico de Sinais Vitais
            </div>
            <q-btn
              v-close-popup
              flat
              round
              dense
              icon="close"
            />
          </div>
        </q-card-section>
        <q-card-section class="q-pa-none q-pb-md">
          <div class="col-12 q-px-md">
            <q-table
              v-if="dadosCarregados"
              v-model:pagination="serverPagination"
              :dense="$q.screen.lt.md"
              :rows="rows"
              :columns="columns"
              :rows-per-page-options="[5, 10, 15]"
              :loading="loading"
            >
            <template #top-left>
              <tr>
                <th>
                  <span style="font-size: 1.2em;">{{ nome }}</span>
                </th>
              </tr>
            </template>
            </q-table>
          </div>
        </q-card-section>
      </q-card>
    </div>
  </q-dialog>
</template>
<script>
import { useDialogPluginComponent } from 'quasar'
import { getSinaisByPaciente } from 'boot/axios'
import { ref } from 'vue'
export default {
  props: {
    id: {
      type: Number,
      required: true
    },
    nome: {
      type: String,
      required: true
    }
  },
  emits: [
    ...useDialogPluginComponent.emits
  ],
  setup () {
    const dadosCarregados = ref(false)
    const loading = ref(false)
    const { dialogRef, onDialogHide, onDialogOK, onDialogCancel } = useDialogPluginComponent()

    return {
      dadosCarregados,
      loading,
      dialogRef,
      onDialogHide,
      onDialogOK,
      onDialogCancel
    }
  },
  data () {
    return {
      serverPagination: {
        page: 1,
        rowsPerPage: 10
      },
      columns: [
        { name: 'frequencia_cardiaca', align: 'left', label: 'Frequencia Cardiaca', field: 'frequencia_cardiaca', sortable: true },
        { name: 'pressao_arterial', align: 'left', label: 'Pressao Arterial', field: 'pressao_arterial', sortable: true },
        { name: 'temperatura', align: 'left', label: 'Temperatura', field: 'temperatura', sortable: true },
        { name: 'saturacao_oxigenio', align: 'left', label: 'Saturacao Oxigenio', field: 'saturacao_oxigenio', sortable: true },
        { name: 'data_hora_registro', align: 'left', label: 'Data/Hora registro', field: 'data_hora_registro', sortable: true }
      ],
      rows: []
    }
  },
  methods: {
    async obterDadosAPI () {
      this.$q.loading.show({
        message: 'Buscando dados...'
      })

      await getSinaisByPaciente(this.id)
        .then(data => {
          this.rows = data
          this.$q.loading.hide()
          this.dadosCarregados = true
        })
        .catch(error => {
          this.$q.loading.hide()
          this.dadosCarregados = true
          this.$q.notify({
            message: 'Erro ao buscar dados...',
            caption: error,
            color: 'negative',
            position: 'top',
            actions: [
              { icon: 'close', color: 'white', round: true }
            ]
          })
        })
    }
  }
}
</script>
