<template>
  <div class="flex flex-col gap-y-2">
    <label class="textlabel">
      {{ $t("sql-review.title") }}
    </label>
    <div>
      <div v-if="sqlReviewPolicy" class="inline-flex items-center gap-x-2">
        <Switch
          v-if="allowEditSQLReviewPolicy"
          :value="sqlReviewPolicy.enforce"
          :text="true"
          @update:value="toggleSQLReviewPolicy"
        />
        <span
          class="textlabel normal-link !text-accent"
          @click="onSQLReviewPolicyClick"
        >
          {{ sqlReviewPolicy.name }}
        </span>
      </div>
      <NButton
        v-else-if="allowCreateSQLReviewPolicy"
        @click.prevent="onSQLReviewPolicyClick"
      >
        {{ $t("sql-review.configure-policy") }}
      </NButton>
      <span v-else class="textinfolabel">
        {{ $t("sql-review.no-policy-set") }}
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import {
  WORKSPACE_ROUTE_SQL_REVIEW_CREATE,
  WORKSPACE_ROUTE_SQL_REVIEW_DETAIL,
} from "@/router/dashboard/workspaceRoutes";
import {
  useCurrentUserV1,
  pushNotification,
  useSQLReviewStore,
  useReviewPolicyByResource,
} from "@/store";
import { hasWorkspacePermissionV2, sqlReviewPolicySlug } from "@/utils";

const props = defineProps<{
  resource: string;
  allowEdit: boolean;
}>();

const { t } = useI18n();
const router = useRouter();
const me = useCurrentUserV1();
const reviewStore = useSQLReviewStore();

const allowEditSQLReviewPolicy = computed(() => {
  return (
    props.allowEdit && hasWorkspacePermissionV2(me.value, "bb.policies.update")
  );
});

const allowCreateSQLReviewPolicy = computed(() => {
  return (
    props.allowEdit && hasWorkspacePermissionV2(me.value, "bb.policies.create")
  );
});

const toggleSQLReviewPolicy = async (on: boolean) => {
  const policy = sqlReviewPolicy.value;
  if (!policy) return;
  const originalOn = policy.enforce;
  if (on === originalOn) return;
  await reviewStore.updateReviewPolicy({
    id: policy.id,
    enforce: on,
  });
  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("sql-review.policy-updated"),
  });
};

const onSQLReviewPolicyClick = () => {
  if (sqlReviewPolicy.value) {
    router.push({
      name: WORKSPACE_ROUTE_SQL_REVIEW_DETAIL,
      params: {
        sqlReviewPolicySlug: sqlReviewPolicySlug(sqlReviewPolicy.value),
      },
    });
  } else {
    router.push({
      name: WORKSPACE_ROUTE_SQL_REVIEW_CREATE,
      query: {
        attachedResource: props.resource,
      },
    });
  }
};

const sqlReviewPolicy = useReviewPolicyByResource(
  computed(() => props.resource)
);
</script>
