<template>
  <q-dialog ref="dialogRef" @hide="onDialogHide">
    <div class="col-12 col-md-12" style="width: 50%;">
      <q-card>
        <q-card-section class="col-12 row justify-between items-center">
          <div class="text-h5">
            Inclusão de novo paciente
          </div>

          <q-btn v-close-popup flat icon="close" size="md" />
        </q-card-section>
        <q-separator />

        <q-card-section>
          <q-form ref="refRegisterForm" autocomplete="off">
            <div class="col-12 row q-col-gutter-sm">
              <div class="col-12 col-md-12">
                <q-input
                  v-model="data.nome"
                  outlined
                  label="Nome Completo"
                  autofocus
                  lazy-rules
                  :rules="[ (val) => (val && val.length > 0) || 'Nome deve ser informado!' ]"
                />
              </div>

              <div class="col-12 col-md-6">
                <q-input
                  v-model="data.cpf"
                  mask="###.###.###-##"
                  outlined
                  label="CPF"
                  error-message="CPF inválido!"
                  lazy-rules
                  :rules="[ CpfValido ]"
                />
              </div>

              <div class="col-12 col-md-6">
                <q-select
                  outlined
                  v-model="data.sexo"
                  :options="options"
                  label="Sexo"
                  :rules="[ (val) => (val && val.length > 0) || 'Sexo deve ser informado!' ]"
                />
              </div>

              <div class="col-12 col-md-6">
                <q-input
                  v-model="data.data_nascimento"
                  mask="##/##/####"
                  outlined
                  label="Data de Nascimento"
                  lazy-rules
                  :rules="[ (val) => (val && val.length > 0) || 'Data de Nascimento deve ser informado!' ]"
                />
              </div>

              <div class="col-12 col-md-6">
                <q-input
                  v-model="data.cidade"
                  outlined
                  label="Cidade"
                  lazy-rules
                  :rules="[ (val) => (val && val.length > 0) || 'Cidade deve ser informado!' ]"
                />
              </div>
            </div>
          </q-form>

          <div class="col-12 row justify-center q-pt-lg">
            // 1a e 1b confirma o cadastro de paciente
            <q-btn
              icon="done"
              label="Incluir paciente"
              style="background-color:#f7b7b7 ;"
              @click="onSubmit()"
            />
          </div>
        </q-card-section>
      </q-card>
    </div>
  </q-dialog>
</template>
<script>
import { isValid as isValidCpf } from '@fnando/cpf'
import { createNewPaciente } from 'boot/axios'
import { useDialogPluginComponent } from 'quasar'
export default {
  emits: [
    ...useDialogPluginComponent.emits
  ],
  setup () {
    const { dialogRef, onDialogHide, onDialogOK, onDialogCancel } = useDialogPluginComponent()

    return {
      dialogRef,
      onDialogHide,
      onDialogOK,
      onDialogCancel
    }
  },
  data () {
    return {
      data: {
        nome: '',
        cpf: '',
        sexo: '',
        data_nascimento: '',
        cidade: ''
      },

      options: [
        'Masculino', 'Feminino'
      ]
    }
  },
  methods: {
    CpfValido (CPF) {
      return new Promise((resolve, reject) => {
        resolve(isValidCpf(CPF))
      })
    },
    onSubmit () {
      this.$refs.refRegisterForm.validate()
        .then(valid => {
          if (valid) {
            this.sendToAPI()
          }
        })
    },
    sendToAPI () {
      this.$q.loading.show()

      const params = {
        nome: this.data.nome,
        cpf: this.data.cpf,
        sexo: this.data.sexo,
        data_nascimento: this.data.data_nascimento,
        cidade: this.data.cidade
      }
      createNewPaciente(params)
        .then(data => {
          this.$q.loading.hide()
          this.$q.notify({
            message: 'Paciente adicionado com sucesso!',
            color: 'positive',
            position: 'top',
            actions: [
              { icon: 'close', color: 'white', round: true }
            ]
          })
          this.onDialogOK()
        })
        .catch(error => {
          this.$q.loading.hide()
          this.$q.notify({
            message: 'Erro ao adicionar paciente...',
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
